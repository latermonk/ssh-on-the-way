<template>
  <div class="console" id="terminal-container"></div>
</template>

<script>
  import {Terminal} from 'xterm';
  import {WebLinksAddon} from 'xterm-addon-web-links';
  import {AttachAddon} from 'xterm-addon-attach';
  import {FitAddon} from 'xterm-addon-fit'
  import '../../node_modules/xterm/css/xterm.css'

  export default {
    name: "WebTerminal",
    mounted() {
      let terminalContainer = document.getElementById('terminal-container');
      let term = new Terminal({
        cols: 80,
        rows: 30,
        cursorBlink: true, // 光标闪烁
        cursorStyle: "underline", // 光标样式  null | 'block' | 'underline' | 'bar'
        scrollback: 80, //回滚
        tabStopWidth: 8, //制表宽度
        screenKeys: true//
      });

      term.loadAddon(new WebLinksAddon());

      let user = this.$route.params.user;
      let pwd = this.$route.params.pwd;
      let port = this.$route.params.port;

      let wsurl = "ws://127.0.0.1:8080/api/v1/webssh" + "?deviceId=demo" + "&user=" + user + "&pwd=" + pwd + "&port=" + port;
      this.ws = new WebSocket(wsurl);
      this.ws.onclose = function (e) {
        alert('连接断开，请刷新页面');
      }

      const attachAddon = new AttachAddon(this.ws);
      term.loadAddon(attachAddon);

      const fitAddon = new FitAddon();
      term.loadAddon(fitAddon);

      term.open(terminalContainer);

      fitAddon.fit();
    }
  }
</script>

<style scoped>
  #terminal-container {
    width: 800px;
    height: 450px;
    margin: 0 auto;
    padding: 2px;
  }
</style>
