package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

var (
	version       = "unknown"
	buildDate     = "2006/01/02-15:04:05"
	gitCommit     = "unknown"
	buildDateTime time.Time
)

func putUsers(ctx context.Context, repo Repo[*user]) {
	repo.Create(ctx, &user{
		name: "Maurizio",
	})
	repo.Create(ctx, &user{
		name: "Irene",
	})
}

func putDogs(ctx context.Context, repo Repo[*dog]) {
	repo.Create(ctx, &dog{
		name:  "Kyrie",
		breed: "Belgian Sheperd",
	})
	repo.Create(ctx, &dog{
		name:  "Ombra",
		breed: "Belgian Sheperd",
	})
}

func printRes[T IDer](ctx context.Context, repo Repo[T]) {
	PrintAll(repo)
	log.Println(repo.GetByID(ctx, 1))
	log.Println(repo.GetByID(ctx, 0))
	log.Println(repo.Remove(ctx, 0))
	log.Println(repo.GetByID(ctx, 0))
	PrintAll(repo)
	log.Println("-------")
}

func userMapRepo(ctx context.Context) {
	userRepo := newMapRepo[*user]()
	putUsers(ctx, userRepo)
	log.Println("Users Map")
	printRes[*user](ctx, userRepo)

}

func userSliceRepo(ctx context.Context) {
	userRepo := newSliceRepo[*user]()
	putUsers(ctx, userRepo)
	log.Println("Users Slice")
	printRes[*user](ctx, userRepo)
}

func dogMapRepo(ctx context.Context) {
	dogRepo := newMapRepo[*dog]()
	putDogs(ctx, dogRepo)
	log.Println("Dogs Map")
	printRes[*dog](ctx, dogRepo)
}

func dogSliceRepo(ctx context.Context) {
	dogRepo := newSliceRepo[*dog]()
	putDogs(ctx, dogRepo)
	log.Println("Dogs Slice")
	printRes[*dog](ctx, dogRepo)
}

func printVersion() {
	fmt.Printf("Version:\t%s\n", version)
	var err error
	buildDateTime, err = time.Parse("2006/01/02-15:04:05", buildDate)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Build date:\t%s\n", buildDateTime)
	fmt.Printf("Git commit:\t%s\n\n\n", gitCommit)
}

func main() {
	ctx := context.TODO()
	printVersion()
	userMapRepo(ctx)
	dogMapRepo(ctx)
	userSliceRepo(ctx)
	dogSliceRepo(ctx)
}
