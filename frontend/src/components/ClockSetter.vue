<template>
    <div :class="containerClass">
      <span :class="[tipsClass, 'marquee']">{{ tips }}</span>
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
      const iconClass = 'icon-[material-symbols--settings] text-blue-300';
      let tips = 'PFinalClub';
      const tipsClass = 'w-5/6 text-black-500 inline-block whitespace-nowrap overflow-hidden';
  
      // Create a reactive reference to props.appConfig
      const appConfig = ref(props.appConfig);
  
      // Watch for changes to props.appConfig
      watch(() => props.appConfig, (newConfig) => {
        appConfig.value = newConfig;
        console.log('Updated appConfig:', newConfig);
      }, { deep: true });
  
      onMounted(() => {
        if (appConfig.value.text_content) {
            tips = appConfig.value.text_content;
        } else {
            tips = 'PFinalClub';
        }
      });
  
      return {
        containerClass,
        iconClass,
        tips,
        tipsClass,
        appConfig
      };
    }
  });
  </script>
  
  <style scoped>
  .marquee {
    overflow: hidden;
    animation: marquee 10s linear infinite;
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
  