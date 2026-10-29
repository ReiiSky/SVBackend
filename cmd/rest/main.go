package main

import backend "github.com/ReiiSky/SharingVisionBackend"

func main() {
	kernel := backend.NewKernel(backend.KernelParams{
		Port: 3000,
	})
	kernel.RunServer()
}
