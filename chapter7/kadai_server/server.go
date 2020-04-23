package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, server.")
}

func fortuneHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, getRandomFortune(r.FormValue("p") == "cheat"))
}

type Request struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

func userFortuneHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストボディの取得
	defer r.Body.Close()
	var req Request
	// io.Readerを実装している
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(req)

	fmt.Fprintf(w, "ID:%dの%sさんの運勢は%sです！", req.UserID, req.Name, getRandomFortune(false))
}

func getRandomFortune(isCheat bool) string {
	if !isCheat {
		switch rand.Intn(4) {
		case 0:
			return "凶"
		case 1:
			return "吉"
		case 2:
			return "中吉"
		}
	}
	return "大吉"
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/fortune", fortuneHandler)
	mux.HandleFunc("/user_fortune", userFortuneHandler)

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// OSからのシグナルを待つ
	go func() {
		// SIGTERM: コンテナが終了する時に送信されるシグナル
		// SIGINT: Ctrl+c
		sigCh := make(chan os.Signal, 1)
		// 受け取るシグナルを指定
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		// チャネルでの待受、シグナルを受け取るまで以降は処理されない
		<-sigCh

		log.Println("start graceful shutdown server.")
		// タイムアウトのコンテキストを設定（後述）
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		// Graceful shutdown
		if err := srv.Shutdown(ctx); err != nil {
			log.Println(err)
			// 接続されたままのコネクションも明示的に切る
			srv.Close()
		}
		log.Println("HTTPServer shutdown.")
	}()

	if err := srv.ListenAndServe(); err != nil {
		log.Print(err)
	}
}
