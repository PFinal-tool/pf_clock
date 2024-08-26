<template>
  <div id="app" @dblclick="toggleTimer" @keydown="handleKeydown" tabindex="0" >
    <FlipClock ref="flipClock"></FlipClock>
  </div>
</template>

<script>
import FlipClock from './components/FlipClock.vue';

export default {
  name: 'app',
  components: {
    FlipClock
  },
  data() {
    return {
      lastKey: ''
    }
  },
  mounted() {
    // 使 div 可以接收焦点
    this.$el.focus()
    // 全局监听键盘事件
    window.addEventListener('keydown', this.handleKeydown)
  },
  beforeDestroy() {
    // 移除全局事件监听
    window.removeEventListener('keydown', this.handleKeydown)
  },
  methods: {
    start() {
      // 确保引用存在后再调用方法
      if (this.$refs.flipClock) {
        this.$refs.flipClock.start()
      } else {
        console.error('FlipClock component is not yet mounted.')
      }
    },
    resetFlipClock() {
      if (this.$refs.flipClock) {
        this.$refs.flipClock.reset()
      }
    },
    handleKeydown(event) {
      // 监听 "ss" 按键
      if (event.key === 's') {
        this.start()
      }else if (event.key === 'p'){
        this.$refs.flipClock.pause()
      }else if (event.key === 'r'){
        this.resetFlipClock()
      }
      // 更新最后按下的键
      this.lastKey = event.key
    }
  }
}
</script>