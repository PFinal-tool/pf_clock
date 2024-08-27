<template>
  <div id="app" @dblclick="toggleTimer" @keydown="handleKeydown" tabindex="0">
    <FlipClock ref="flipClock"></FlipClock>
    <ClockSetter :appConfig="appConfig"></ClockSetter>
  </div>
</template>

<script>
import { onBeforeUnmount, onMounted, reactive, ref } from 'vue';
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

    // 获取配置信息
    GetAppConfig().then(config => {
      Object.assign(appConfig, config);
    }).catch(error => {
      console.error('Error fetching config:', error);
    });

    const start = () => {
      if (flipClockRef.value) {
        flipClockRef.value.start();
      } else {
        console.error('FlipClock component is not yet mounted.');
      }
    };

    const resetFlipClock = () => {
      if (flipClockRef.value) {
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

    const flipClockRef = ref(null);

    onMounted(() => {
      document.getElementById('app').focus();
      window.addEventListener('keydown', handleKeydown);
    });

    onBeforeUnmount(() => {
      window.removeEventListener('keydown', handleKeydown);
    });

    return {
      appConfig,
      lastKey,
      flipClockRef,
    };
  }
}
</script>
