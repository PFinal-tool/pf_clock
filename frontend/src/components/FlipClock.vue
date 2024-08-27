<template>
  <div class="FlipClock">
    <Flipper ref="flipperHour1" />
    <Flipper ref="flipperHour2" />
    <em>:</em>
    <Flipper ref="flipperMinute1" />
    <Flipper ref="flipperMinute2" />
    <em>:</em>
    <Flipper ref="flipperSecond1" />
    <Flipper ref="flipperSecond2" />
  </div>

</template>

<script>
import Flipper from './Flipper.vue'

export default {
  name: 'FlipClock',
  data() {
    return {
      timer: null,
      flipObjs: [],
      isRunning: false,
      startTime: 0, // 用于存储计时器开始的时间戳
    }
  },
  components: {
    Flipper
  },
  methods: {
    // 初始化数字
    init() {
      let nowTimeStr = "000000"
      console.log(nowTimeStr)
      for (let i = 0; i < this.flipObjs.length; i++) {
        this.flipObjs[i].setFront(nowTimeStr[i])
      }
    },
    // 开始计时
    start() {
      if (this.isRunning) {
        return
      }
      this.isRunning = true
      this.run()
    },
    // 暂停计时
    pause() {
      if (!this.isRunning) {
        return
      }
      this.isRunning = false
      clearInterval(this.timer)
    },
    // 重置计时
    reset() {
      this.isRunning = false
      clearInterval(this.timer)
      this.startTime = 0
      this.init()
    },
    // 
    run() {
      this.timer = setInterval(() => {
        let nowTimeStr, nextTimeStr;

        if (this.startTime > 0) {
          // Format current and next time
          nowTimeStr = this.formatDate(this.startTime, 'hhiiss');
          nextTimeStr = this.formatDate(this.startTime + 1, 'hhiiss');
        } else {
          // Initialize to "000000" and "000001" for the first run
          nowTimeStr = "000000";
          nextTimeStr = "000001";
        }
        for (let i = 0; i < this.flipObjs.length; i++) {
          if (nowTimeStr[i] === nextTimeStr[i]) {
            continue
          }
          this.flipObjs[i].flipDown(
            nowTimeStr[i],
            nextTimeStr[i]
          )
        }
        this.startTime++
      }, 1000)
    },
    // 正则格式化日期
    formatDate(seconds, format) {
      // 将秒数转换为小时、分钟和秒数
      let hours = Math.floor(seconds / 3600) % 24;
      let minutes = Math.floor(seconds / 60) % 60;
      let secs = seconds % 60;

      // 构建格式化后的时间字符串
      let formatted = format
        .replace(/h+/g, match => ('0' + hours).slice(-match.length))
        .replace(/i+/g, match => ('0' + minutes).slice(-match.length))
        .replace(/s+/g, match => ('0' + secs).slice(-match.length));

      return formatted;
    },
    // 日期时间补零
    padLeftZero(str) {
      return ('00' + str).substr(str.length)
    }
  },
  mounted() {
    this.flipObjs = [
      this.$refs.flipperHour1,
      this.$refs.flipperHour2,
      this.$refs.flipperMinute1,
      this.$refs.flipperMinute2,
      this.$refs.flipperSecond1,
      this.$refs.flipperSecond2
    ]
    this.init()
    // this.run()
  }
}
</script>

<style>
.FlipClock {
  text-align: center;
  padding-top: 1em;
}

.FlipClock .M-Flipper {
  margin: 0 3px;
}

.FlipClock em {
  display: inline-block;
  line-height: 102px;
  font-size: 66px;
  font-style: normal;
  vertical-align: top;
}
</style>