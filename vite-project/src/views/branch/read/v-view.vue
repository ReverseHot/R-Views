<template>

    <div v-if="store.loader_content">
        <v-skeleton-loader type="avatar,text,button"></v-skeleton-loader>
        <v-skeleton-loader type="avatar,text,button"></v-skeleton-loader>
        <v-skeleton-loader type="avatar,text,button"></v-skeleton-loader>
    </div>

    <div v-else>
        <v-banner color="#26A69A" icon="mdi-web" lines="one">
            <v-banner-text>
                访问地址
            </v-banner-text>

            <template v-slot:actions>
                <v-btn @click="copyToUrlWeb">复制</v-btn>
            </template>
        </v-banner>

        <v-banner color="error" icon="mdi-bash" lines="one">
            <v-banner-text>
                接口地址
            </v-banner-text>

            <template v-slot:actions>
                <v-btn @click="copyToUrlApi">复制</v-btn>
            </template>
        </v-banner>

        <v-banner color="#388E3C" icon="mdi-content-copy" lines="one">
            <v-banner-text>
                直接复制
            </v-banner-text>

            <template v-slot:actions>
                <v-btn @click="copyToUrlText">复制</v-btn>
            </template>
        </v-banner>
    </div>

    <div>
        <v-snackbar v-model="message" variant="text">
            <v-alert icon="mdi-check" text="复制成功！" type="success"></v-alert>
        </v-snackbar>
    </div>

</template>

<script setup>
import { ref } from 'vue'
import { useStoreBranchRead } from "@/store/store_branch_read"

const store = useStoreBranchRead()
const message = ref(false)


// 访问地址
const copyToUrlWeb = async () => {
    try {
        // await navigator.clipboard.writeText('要复制的文本内容');
        await navigator.clipboard.writeText(store.respond_comments.url.web)
        message.value = true;
    } catch (error) {
        console.log("复制失败:", error);
    }
}


// 接口地址
const copyToUrlApi = async () => {
    try {
        await navigator.clipboard.writeText(store.respond_comments.url.api)
        message.value = true;
    } catch (error) {
        console.log("复制失败:", error);
    }
}


// 原始内容
const copyToUrlText = async () => {
    try {
        await navigator.clipboard.writeText(store.respond_content)
        message.value = true;
    } catch (error) {
        console.log("复制失败:", error);
    }
}

</script>
