<template>
  <div id="app" @dblclick="toggleTimer" @keydown="handleKeydown" tabindex="0">
    <FlipClock ref="flipClockRef"></FlipClock>
    <ClockSetter :appConfig="appConfig"></ClockSetter>
  </div>
</template>

<script>
import { nextTick, onBeforeUnmount, onMounted, reactive, ref } from 'vue';
import { GetAppConfig } from '../wailsjs/go/main/App';
import ClockSetter from './components/ClockSetter.vue';
import FlipClock from './components/FlipClock.vue';

export default {
  name: 'App',
  components: {
    FlipClock,
    ClockSetter
  },
  setup() {
    const appConfig = reactive({});
    const lastKey = ref('');
    const flipClockRef = ref(null);
    const isFlipClockMounted = ref(false); // Flag to check if FlipClock is mounted

    // 获取配置信息
    GetAppConfig().then(config => {
      Object.assign(appConfig, config);
    }).catch(error => {
      console.error('Error fetching config:', error);
    });

    const start = () => {
      if (flipClockRef.value && isFlipClockMounted.value) {
        console.log('FlipClock is mounted and ready to start');
        flipClockRef.value.start();
      } else {
        console.error('FlipClock component is not yet mounted.');
        console.log('flipClockRef.value:', flipClockRef.value);
      }
    };

    const resetFlipClock = () => {
      if (flipClockRef.value && isFlipClockMounted.value) {
        console.log('Resetting FlipClock');
        flipClockRef.value.reset();
      }
    };

    const handleKeydown = (event) => {
      if (event.key === 's') {
        start();
      } else if (event.key === 'p') {
        flipClockRef.value.pause();
      } else if (event.key === 'r') {
        resetFlipClock();
      }
      lastKey.value = event.key;
    };

    onMounted(async () => {
      await nextTick(); // Wait for the DOM to update
      console.log(flipClockRef.value)
      document.getElementById('app').focus();
      window.addEventListener('keydown', handleKeydown);
      isFlipClockMounted.value = true;
      console.log('App mounted');
    });

    onBeforeUnmount(() => {
      window.removeEventListener('keydown', handleKeydown);
    });

    return {
      appConfig,
      lastKey,
      flipClockRef,
      start,
      resetFlipClock
    };
  }
}
</script>
