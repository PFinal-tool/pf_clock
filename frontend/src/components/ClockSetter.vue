<template>
  <div :class="containerClass" :style="computedStyle">
    <span :class="[tipsClass, 'marquee']">
      <h5>{{ tips }}</h5>
    </span>
    <span :class="iconClass"></span>
  </div>
</template>

<script>
import { defineComponent, onMounted, ref, watch } from 'vue';

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
    const computedStyle = ref('');
    // Create a reactive reference to props.appConfig
    const appConfig = ref(props.appConfig);

    watch(() => props.appConfig, (newConfig) => {
       if (newConfig) {
          appConfig.value = newConfig;
        }
        if (props.appConfig && props.appConfig.text_content) {
          tips.value = props.appConfig.text_content;
        }
        const style = {};
        if (props.appConfig && props.appConfig.text_color) {
          style.color = props.appConfig.text_color;
        }
        if (props.appConfig && props.appConfig.text_bg_color) {
          style.backgroundColor = props.appConfig.text_bg_color;
        }
        computedStyle.value = style;
    }, { deep: true });

    onMounted(() => {
      try {
        // Set initial value based on appConfig
        if (props.appConfig && props.appConfig.text_content) {
          tips.value = props.appConfig.text_content;
        }
        const style = {};
        if (props.appConfig && props.appConfig.text_color) {
          style.color = props.appConfig.text_color;
        }
        if (props.appConfig && props.appConfig.text_bg_color) {
          style.backgroundColor = props.appConfig.text_bg_color;
        }
        computedStyle.value = style;
        console.log(computedStyle)
      } catch (error) {
        console.error('Error in onMounted callback:', error);
      }
    });

    return {
      containerClass,
      iconClass,
      tips,
      tipsClass,
      appConfig,
      computedStyle
    };
  }
});
</script>

<style scoped>
.marquee {
  overflow: hidden;
}
.marquee h5 {
  animation: marquee 8s linear infinite;
}

@keyframes marquee {
  from {
    transform: translateX(100%);
  }

  to {
    transform: translateX(-100%);
  }
}
</style>