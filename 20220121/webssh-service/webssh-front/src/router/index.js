import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import WebTerminal from "../components/WebTerminal";

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/terminal',
      name: 'WebTerminal',
      component: WebTerminal
    }
  ]
})
