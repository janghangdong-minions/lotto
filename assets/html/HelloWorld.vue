<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <h2>Essential Links</h2>
    <ul>
      <li>
        <a
          href="https://vuejs.org"
          target="_blank"
        >
          Core Docs
        </a>
      </li>
      <li>
        <a
          href="https://forum.vuejs.org"
          target="_blank"
        >
          Forum
        </a>
      </li>
      <li>
        <a
          href="https://chat.vuejs.org"
          target="_blank"
        >
          Community Chat
        </a>
      </li>
      <li>
        <a
          href="https://twitter.com/vuejs"
          target="_blank"
        >
          Twitter
        </a>
      </li>
      <br>
      <li>
        <a
          href="http://vuejs-templates.github.io/webpack/"
          target="_blank"
        >
          Docs for This Template
        </a>
      </li>
    </ul>
    <h2>Ecosystem</h2>
    <ul>
      <li>
        <a
          href="http://router.vuejs.org/"
          target="_blank"
        >
          vue-router
        </a>
      </li>
      <li>
        <a
          href="http://vuex.vuejs.org/"
          target="_blank"
        >
          vuex
        </a>
      </li>
      <li>
        <a
          href="http://vue-loader.vuejs.org/"
          target="_blank"
        >
          vue-loader
        </a>
      </li>
      <li>
        <a
          href="https://github.com/vuejs/awesome-vue"
          target="_blank"
        >
          awesome-vue
        </a>
      </li>
    </ul>
    <ul>
      <div class="py-5">
        <h1>로또 추첨기</h1>
        <div>
          {{ message }}
        </div>
      </div>
      <div id=main>
        <div v-for="idx in lottos">
          <span>{{ idx.num }}, {{ idx.color }}</span>
        </div>
      </div>
      <div class="py-3">
        <div id="button_group">
          <button @click="getRandomNums()">새로뽑기</button>
        </div>
      </div>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App',
      lottos: []
    }
  },
  methods: {
    getRandomNums() {
      const lotto = []
      for (let i = 0; i < 7; i++) {
        let color = ''
        let num = Math.floor(Math.random() * 44) + 1
        if (num <= 10) {
          color = 'yellow'
        } else if (num > 10 && num <= 20) {
          color = 'blue'
        } else if (num > 20 && num <= 30) {
          color = 'red'
        } else if (num > 30 && num <= 40) {
          color = 'black'
        } else if (num > 40 && num <= 50) {
          color = 'green'
        }
        num = String(num)

        lotto[i] = { color, num }
      }
      this.lottos = lotto
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
