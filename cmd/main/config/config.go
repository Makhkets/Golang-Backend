package config

import "gopkg.in/go-playground/validator.v9"

var Validate *validator.Validate

const (
	Port      string = "8000"
	MainTable string = "tasks"

	// colors
	Red    string = "\033[31m"
	Green  string = "\033[32m"
	Yellow string = "\033[33m"
	Blue   string = "\033[34m"
	Purple string = "\033[35m"
	Cyan   string = "\033[36m"
	Reset  string = "\033[0m"
)

var Scheme []string = []string{
	`
	CREATE TABLE tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		author TEXT NOT NULL
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
