<template>
    <div :class="containerClass">
        <span :class="[tipsClass,'marquee']">{{ tips }}</span>
        <span :class="iconClass"></span> 
    </div>
</template>

<script>
export default {
    name: 'ClockSetter',
    data() {
        return {
            containerClass: 'w-11/12 h-10 bg-gray-200 rounded-lg flex justify-center items-center m-auto',
            iconClass: 'icon-[material-symbols--settings] text-blue-300',
            tips: "PFinalClub",
            tipsClass: "w-5/6 text-black-500 inline-block whitespace-nowrap"
        }
    }
}
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
/* 如果有需要，可以在此添加样式 */
</style>
