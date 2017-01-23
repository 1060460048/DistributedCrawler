/*
 * 调度器(Scheduler) for master
 * 用来接受引擎发过来的请求, 压入队列中, 并在引擎再次请求的时候返回.
 * 可以想像成一个URL（抓取网页的网址或者说是链接）的优先队列, 由它来决定下一个要抓取的网址是什么, 同时去除重复的网址
 */
 package scrawler

 import (
   "fmt"
 )

 type Item struct {
   name string
   sex  string
   habbit []string
   //define by yourself...
 }

 func Scheduler(cookie string, urls []string) {
   fmt.Printf("Scheduler======")
   //do something for urls maybe use some threads go func() to opti
   for _, url := range urls {
     go Downloader(cookie, url)
   }
 }
