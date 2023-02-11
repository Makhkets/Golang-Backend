package main

const (
	Port string = "8000"

	// colors
	red    string = "\033[31m"
	green  string = "\033[32m"
	yellow string = "\033[33m"
	blue   string = "\033[34m"
	purple string = "\033[35m"
	cyan   string = "\033[36m"
	reset  string = "\033[0m"
)

var Scheme []string = []string{
	`
					CREATE TABLE users (
						id INTEGER PRIMARY KEY AUTOINCREMENT,
						username TEXT NOT NULL,
						password TEXT NOT NULL,
						is_admin BOOLEAN NOT NULL DEFAULT 0
					);
	`,
}

// Example creating tables

//var Scheme []string = []string{
//	`
//					CREATE TABLE users (
//						id INTEGER PRIMARY KEY AUTOINCREMENT,
//						username TEXT NOT NULL,
//						password TEXT NOT NULL,
//						is_admin BOOLEAN NOT NULL DEFAULT 0
//					);
//	`,
//	`
//					CREATE TABLE admins (
//						id INTEGER PRIMARY KEY AUTOINCREMENT,
//						username TEXT NOT NULL,
//						password TEXT NOT NULL,
//						is_admin BOOLEAN NOT NULL DEFAULT 0
//					);
//	`,
//	`
//					CREATE TABLE tests (
//						id INTEGER PRIMARY KEY AUTOINCREMENT,
//						username TEXT NOT NULL,
//						password TEXT NOT NULL,
//						is_admin BOOLEAN NOT NULL DEFAULT 0
//					);
//	`,
//}
