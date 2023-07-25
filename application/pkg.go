package application

import (
	"base-go/application/auth"
	"base-go/application/blogs"
	"base-go/application/cats"
	"base-go/application/events"
	"base-go/application/users"
	"base-go/application/vouchers"
)

type App struct {
	Auth     auth.AuthInteractor
	Cats     cats.CatsInteractor
	Users    users.UsersInteractor
	Blogs    blogs.BlogsInteractor
	Events   events.EventsInteractor
	Vouchers vouchers.VouchersInteractor
}

func NewApp(auth auth.AuthInteractor, cats cats.CatsInteractor, users users.UsersInteractor, blogs blogs.BlogsInteractor, events events.EventsInteractor, vouchers vouchers.VouchersInteractor) *App {
	return &App{
		Auth:     auth,
		Cats:     cats,
		Users:    users,
		Blogs:    blogs,
		Events:   events,
		Vouchers: vouchers,
	}
}
