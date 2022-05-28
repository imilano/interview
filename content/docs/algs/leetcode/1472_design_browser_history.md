---
title: 1472. Design Browser History
weight: 10
tags: [
	"Stack",
]
---

## Description
> You have a browser of one tab where you start on the homepage and you can visit another url, get back in the history number of steps or move forward in the history number of steps.
> 
> Implement the BrowserHistory class:
> 
> - BrowserHistory(string homepage) Initializes the object with the homepage of the browser.
> - void visit(string url) Visits url from the current page. It clears up all the forward history.
> - string back(int steps) Move steps back in history. If you can only return x steps in the history and steps > x, you will return only x steps. Return the current url after moving back in history at most steps.
> - string forward(int steps) Move steps forward in history. If you can only forward x steps in the history and steps > x, you will forward only x steps. Return the current url after forwarding in history at most steps.

## Solutions

简单题，维护一个栈即可。
```go
type BrowserHistory struct {
    urls []string
    pos int
}


func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{urls: []string{homepage}, pos: 0}
}


func (this *BrowserHistory) Visit(url string)  {
    (*this).urls = (*this).urls[:(*this).pos+1]
    (*this).urls = append((*this).urls, url)
    (*this).pos++
}


func (this *BrowserHistory) Back(steps int) string {
    size := (*this).pos+1
    if size <= steps {
        ele := (*this).urls[0]
        (*this).pos = 0
        return ele
    } 
    
    (*this).pos -= steps
    return (*this).urls[(*this).pos]
}


func (this *BrowserHistory) Forward(steps int) string {
    size := len((*this).urls[(*this).pos+1:])
    if size <= steps {
        (*this).pos = len((*this).urls)-1
        return (*this).urls[(*this).pos]
    }
    
    (*this).pos += steps
    return (*this).urls[(*this).pos]
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
 ```
