package main

type subject interface{
	register(Observer)
	deregister(Observer)
	notifyAll()
} 