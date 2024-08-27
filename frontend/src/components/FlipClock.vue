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

<script setup>
import { nextTick, onMounted, ref } from 'vue';
import Flipper from './Flipper.vue';

// Declare refs
const flipperHour1 = ref(null);
const flipperHour2 = ref(null);
const flipperMinute1 = ref(null);
const flipperMinute2 = ref(null);
const flipperSecond1 = ref(null);
const flipperSecond2 = ref(null);

const flipObjs = ref([]); // Reactive reference to flipObjs
const timer = ref(null);
const isRunning = ref(false);
const startTime = ref(0); // Used to store the timer start timestamp

// Initialize the flip clock
function init() {
  let nowTimeStr = "000000";
  console.log(nowTimeStr);
  flipObjs.value.forEach((flipObj, index) => {
    flipObj.setFront(nowTimeStr[index]);
  });
}

// Start the timer
function start() {
  if (isRunning.value) {
    return;
  }
  isRunning.value = true;
  run();
}

// Pause the timer
function pause() {
  if (!isRunning.value) {
    return;
  }
  isRunning.value = false;
  clearInterval(timer.value);
}

// Reset the timer
function reset() {
  isRunning.value = false;
  clearInterval(timer.value);
  startTime.value = 0;
  init();
}

// Run the timer
function run() {
  timer.value = setInterval(() => {
    let nowTimeStr, nextTimeStr;

    if (startTime.value > 0) {
      nowTimeStr = formatDate(startTime.value, 'hhiiss');
      nextTimeStr = formatDate(startTime.value + 1, 'hhiiss');
    } else {
      nowTimeStr = "000000";
      nextTimeStr = "000001";
    }
    flipObjs.value.forEach((flipObj, index) => {
      if (nowTimeStr[index] !== nextTimeStr[index]) {
        flipObj.flipDown(nowTimeStr[index], nextTimeStr[index]);
      }
    });
    startTime.value++;
  }, 1000);
}

// Format date
function formatDate(seconds, format) {
  let hours = Math.floor(seconds / 3600) % 24;
  let minutes = Math.floor(seconds / 60) % 60;
  let secs = seconds % 60;

  return format
    .replace(/h+/g, match => ('0' + hours).slice(-match.length))
    .replace(/i+/g, match => ('0' + minutes).slice(-match.length))
    .replace(/s+/g, match => ('0' + secs).slice(-match.length));
}

// On mounted lifecycle hook
onMounted(async () => {
  await nextTick(); // Ensure DOM is fully updated
  flipObjs.value = [
    flipperHour1.value,
    flipperHour2.value,
    flipperMinute1.value,
    flipperMinute2.value,
    flipperSecond1.value,
    flipperSecond2.value
  ];
  init();
  // Uncomment this line to start the timer automatically
  // run();
});

defineExpose({
  start,
  pause,
  reset
});

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
