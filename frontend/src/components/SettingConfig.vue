<template>
    <div class="min-h-screen flex flex-col items-center justify-center p-6 text-slate-700">
        <div class="w-full max-w-lg p-8">
            <h2 class="text-2xl font-semibold text-white mb-6 text-center">编辑时钟设置</h2>

            <form @submit.prevent="submitForm" class="space-y-6">
                <div class="flex items-center space-x-4">
                    <!-- 时钟背景颜色 -->
                    <div class="flex items-center">
                        <label for="clock_bg_color" class="text-sm font-medium text-white mr-2">时钟背景颜色</label>
                        <input type="color" id="clock_bg_color" v-model="settings.clock_bg_color"
                               class="p-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:outline-none" />
                    </div>

                    <!-- 时钟文本颜色 -->
                    <div class="flex items-center">
                        <label for="clock_text_color" class="text-sm font-medium text-white mr-2">时钟文本颜色</label>
                        <input type="color" id="clock_text_color" v-model="settings.clock_text_color"
                               class="p-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:outline-none" />
                    </div>
                </div>

                <div class="flex items-center space-x-4">
                    <!-- 文本背景颜色 -->
                    <div class="flex items-center">
                        <label for="text_bg_color" class="text-sm font-medium text-white mr-2">文本背景颜色</label>
                        <input type="color" id="text_bg_color" v-model="settings.text_bg_color"
                               class="p-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:outline-none" />
                    </div>

                    <!-- 文本颜色 -->
                    <div class="flex items-center">
                        <label for="text_color" class="text-sm font-medium text-white mr-2">文本颜色</label>
                        <input type="color" id="text_color" v-model="settings.text_color"
                               class="p-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:outline-none" />
                    </div>
                </div>

                <!-- 文本内容 -->
                <div class="flex items-center space-x-4">
                    <label for="text_content" class="text-sm font-medium text-white">文本内容</label>
                    <input type="text" id="text_content" v-model="settings.text_content"
                           class="flex-1 p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:outline-none"
                           placeholder="输入文本内容" />
                </div>

                <!-- 提交按钮 -->
                <div class="text-center flex justify-between">
                    <button type="button"
                            class="w-1/4 bg-blue-500 text-white p-3 rounded-lg hover:bg-blue-600 focus:ring-4 focus:ring-blue-300 focus:outline-none"
                            @click="submitForm">
                        更新设置
                    </button>
                    <button type="button"
                            class="w-1/4 bg-gray-400 text-white p-2 rounded-lg hover:bg-green-400 focus:ring-4 focus:ring-blue-300 focus:outline-none"
                            @click="resetForm">
                        取消设置
                    </button>
                </div>


            </form>

            <!-- 显示 ClockSetter -->
            <div :class="containerClass" :style="computedStyle" class="mt-8">
                <span :class="[tipsClass, 'marquee']">
                    <h5 :style="textStyles">{{ settings.text_content }}</h5>
                </span>
                <span :class="iconClass" @click="settingClick"></span>
            </div>
        </div>
    </div>
</template>

<script>
import { ref, reactive, watch } from "vue";
import { Setting,Reload } from "../../wailsjs/go/main/App";
export default {
    name: "SettingConfig", // Bug 修复：组件名称应该与模板中的组件名称一致
    setup() {
        // JSON 配置
        const settings = reactive({
            clock_bg_color: "#FFFFFF",
            clock_text_color: "#000000",
            text_bg_color: "#1ab394",
            text_color: "red",
            text_content: "PFinalClub专业的技术社区",
        });

        const containerClass =
            "w-11/12 h-10 rounded-lg flex justify-center items-center m-auto overflow-hidden";
        const iconClass = "icon-[material-symbols--settings] text-blue-300 w-1/12";
        const tipsClass = "w-11/12 overflow-hidden";

        const computedStyle = reactive({
            color: settings.clock_text_color,
            backgroundColor: settings.text_bg_color,
        });

        // 用于文字部分的样式
        const textStyles = reactive({
            color: settings.text_color,
        });

        const settingClick = () => {
            console.log("Settings button clicked");
        };

        const submitForm = () => {
            // 更新样式时，将表单中的值应用到样式上
            computedStyle.color = settings.clock_text_color;
            computedStyle.backgroundColor = settings.text_bg_color;
            textStyles.color = settings.text_color;

            console.log("Settings updated:", settings);
        };

        const resetForm = () => {
            // 重置表单时，将样式重置为初始值
            console.log("Settings reset");
            Reload();
        };
        // 使用 watch 监听 settings 的变化
        watch(
            () => settings,
            (newSettings) => {
                computedStyle.color = newSettings.clock_text_color;
                computedStyle.backgroundColor = newSettings.text_bg_color;
                textStyles.color = newSettings.text_color;
            },
            { deep: true }
        );

        return {
            settings,
            containerClass,
            iconClass,
            tipsClass,
            computedStyle,
            textStyles,
            settingClick,
            submitForm,
        };
    },
};
</script>

<style scoped>
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

input[type="color"] {
    padding: 0;
    border: none;
    border-radius: 0%;
    outline: none;
    /*清除input点击之后的黑色边框*/
    box-sizing: border-box;
    /* 包含内边距在宽度内 */
}
</style>
