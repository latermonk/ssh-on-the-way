# webssh-前端

> 基于[xtermjs](https://xtermjs.org/)的在线ssh终端
>
> 后端地址 -> [webssh-service](https://github.com/math-good/webssh-service)

## 配置

websocket地址(src/components/WebTerminal.vue)
```ts
let wsurl = "ws://127.0.0.1:8080/api/v1/webssh" + "?deviceId=demo"
```

## Build Setup

``` bash
# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build

# build for production and view the bundle analyzer report
npm run build --report
```

