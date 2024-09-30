<template>
  <div id="app" @dblclick="toggleTimer" @keydown="handleKeydown" tabindex="0">
    <!-- 控制 FlipClock 的显示 -->
    <FlipClock ref="flipClockRef" v-if="flipClockVisible"></FlipClock>

    <!-- 控制 ClockSetter 的显示 -->
    <ClockSetter
      :appConfig="appConfig"
      @toggle-visibility="handleVisibilityChange"
      v-if="settingStatus"
    ></ClockSetter>

    <!-- 控制 SettingConfig 的显示 -->
    <SettingConfig
      v-if="settingConfigVisible"
      @back-to-clocksetter="handleBackToClockSetter"
    ></SettingConfig>
  </div>
</template>

<script>
import { nextTick, onBeforeUnmount, onMounted, reactive, ref } from "vue";
import { GetAppConfig } from "../wailsjs/go/main/App";
import ClockSetter from "./components/ClockSetter.vue";
import FlipClock from "./components/FlipClock.vue";
import SettingConfig from "./components/SettingConfig.vue"; // 引入 SettingConfig 组件

export default {
  name: "App",
  components: {
    FlipClock,
    ClockSetter,
    SettingConfig, // 注册 SettingConfig 组件
  },
  setup() {
    const appConfig = reactive({});
    const flipClockRef = ref(null);
    const settingStatus = ref(true); // 控制 ClockSetter 显示
    const flipClockVisible = ref(true); // 控制 FlipClock 显示
    const settingConfigVisible = ref(false); // 控制 SettingConfig 显示

    // 获取配置信息
    GetAppConfig()
      .then((config) => {
        Object.assign(appConfig, config);
      })
      .catch((error) => {
        console.error("Error fetching config:", error);
      });

    // 处理 ClockSetter 传递的事件，控制 FlipClock 和 ClockSetter 的可见性
    const handleVisibilityChange = () => {
      settingStatus.value = false; // 隐藏 ClockSetter
      flipClockVisible.value = false; // 隐藏 FlipClock
      settingConfigVisible.value = true; // 显示 SettingConfig
    };

    // 处理 SettingConfig 传递的事件，返回 ClockSetter 界面
    const handleBackToClockSetter = () => {
      settingConfigVisible.value = false; // 隐藏 SettingConfig
      settingStatus.value = false; // 显示 ClockSetter
    };

    const start = () => {
      if (flipClockRef.value) {
        flipClockRef.value.start();
      }
    };

    const resetFlipClock = () => {
      if (flipClockRef.value) {
        flipClockRef.value.reset();
      }
    };

    const handleKeydown = (event) => {
      if (event.key === "s") {
        start();
      } else if (event.key === "p") {
        flipClockRef.value.pause();
      } else if (event.key === "r") {
        resetFlipClock();
      }
    };

    onMounted(async () => {
      await nextTick();
      document.getElementById("app").focus();
      window.addEventListener("keydown", handleKeydown);
    });

    onBeforeUnmount(() => {
      window.removeEventListener("keydown", handleKeydown);
    });

    return {
      appConfig,
      flipClockRef,
      start,
      resetFlipClock,
      settingStatus,
      flipClockVisible,
      settingConfigVisible,
      handleVisibilityChange,
      handleBackToClockSetter,
    };
  },
};
</script>
