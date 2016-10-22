package scrawler

import (
  "fmt"
)

/*
 * open main page and you should get cookie and save
 */
 func getCookies() (strCookies string) {
   strLoginUrl := `http://weibo.com/login.php`
   _, strCookies := DoRequest(`GET`, strLoginUrl, ``, ``, ``, nil)
 }

/*
 * when finish inputing the username, send the prelogin req
 * you can get login info for logining sina
 */
func sendPreLogin
