package main

import "time"

type Post struct {
  Id       int       `json:"id"`
  Content  string    `json:"content"`
  Author   Author    `json:"author"`
  Comments []Comment `json:"comments"`
}

type Author struct {
  Id   int         `json:"id"`
  Name string      `json:"name"`
}

type Comment struct {
  Id      int         `json:"id"`
  Content string      `json:"content"`
  Author  string      `json:"author"`
}


type Scenario struct {
  Id       int       `json:"id"`
  Title    string    `json:"title"`
  Pause    time.Duration       `json:"pause"`
}



