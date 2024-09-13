<template>
  <div :class="containerClass" :style="computedStyle">
    <span :class="[tipsClass, 'marquee']">
      <h5>{{ tips }}</h5>
    </span>
    <span :class="iconClass" @click="settingClick"></span>
  </div>
</template>

<script>
import { defineComponent, onMounted, ref, watch, watchEffect } from 'vue';
import { Setting } from '../../wailsjs/go/main/App';

export default defineComponent({
  name: 'ClockSetter',
  props: {
    appConfig: {
      type: Object,
      required: true
    }
  },
  setup(props) {
    const containerClass = 'w-11/12 h-10 bg-gray-200 rounded-lg flex justify-center items-center m-auto overflow-hidden';
    const iconClass = 'icon-[material-symbols--settings] text-blue-300 w-1/12';
    const tips = ref('PFinalClub');
    const tipsClass = 'w-11/12 overflow-hidden';
    const computedStyle = ref({});

    const settingClick = () => {
      console.log('设置按钮被点击了');
      console.log(Setting())
    };

    watchEffect(() => {
      const config = props.appConfig;
      tips.value = config.text_content || 'PFinalClub';
      computedStyle.value = {
        color: config.text_color || '',
        backgroundColor: config.text_bg_color || ''
      };
    });

    return {
      containerClass,
      iconClass,
      tips,
      tipsClass,
      computedStyle,
      settingClick
    };
  }
});
</script>

<style scoped>
.marquee {
  overflow: hidden;
}

.marquee h5 {
  animation: marquee 4s linear infinite;
}

@keyframes marquee {
  0%,
  100% {
    transform: translateX(-25%);
  }

  50% {
    transform: translateX(25%);
  }
}

@-webkit-keyframes marquee {
  0% {
    -webkit-transform: translateX(100%);
  }
  100% {
    -webkit-transform: translateX(-100%);
  }
}
</style>
