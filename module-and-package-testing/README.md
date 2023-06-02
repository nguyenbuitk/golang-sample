Để có thể sử dụng package1 ở module hoặc package thì phải copy package đó vào $GOROOT/src
    sudo cp -rf ./pacakge /usr/local/go/src

Để có thể sử dụng local module1 ở local module2, thì cần thực hiện các bước sau tại module2:
    # import module 1 (tại main.go)
    import ("github.com/nguyenbuitk/module1")

    # replace module1 to local module1
    go go mod edit -replace=github.com/nguyenbuitk/module1=../module1
    go mod tidy
