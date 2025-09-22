<template>
  <div class="home-view">
    <n-space size="large" spacing="100" vertical>
      <h1>添加计划</h1>
      <!--align 居中对齐-->
      <n-space size="large" spacing="16" align="center">
        <n-tag class="tag">计划标题</n-tag>
        <n-input v-model="value" placeholder="起一个帅一点的标题吧" class="responsive-input" />
      </n-space>
      <!--align 居中对齐-->
      <n-space size="large" spacing="16" align="center">
        <n-tag class="tag">计划内容</n-tag>
        <n-input v-model="value" placeholder="详细描述一下你的想法吧" class="responsive-input" type="textarea" :rows=10 />
      </n-space>
      <n-space size="large" spacing="16" align="center">
        <n-tag class="tag">排个计划</n-tag>
        <n-date-picker v-model="Date" type="date" placeholder="选择日期" />
        <n-radio :checked="checkedValue === 'Definitely Maybe'" value="Definitely Maybe" name="basic-demo"
          @change="handleChange">
          Definitely Maybe
        </n-radio>
        <n-radio :checked="checkedValue === 'Be Here Now'" value="Be Here Now" name="basic-demo" @change="handleChange">
          Be Here Now
        </n-radio>
      </n-space>
    </n-space>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'


const checkedValueRef = ref<string | null>(null)
const disabled = ref(true)
const checkedValue = checkedValueRef

function handleChange(e: Event) {
  checkedValueRef.value = (e.target as HTMLInputElement).value
}

const value = ref<string>('123')
const value1 = ref<string>('12354')
const planContent = ref<string>('这是计划内容')
const Date = ref<string>('')
</script>

<style scoped>
.home-view {
  /* 确保容器占满整个屏幕 */
  position: fixed;
  top: -20%;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 基础样式：占满父容器宽度 */
.responsive-input {
  width: 100%;
  /* 最大宽度限制，避免在超大屏幕上过宽 */
  max-width: 800px;
}

/* 屏幕宽度 >= 1200px（大屏幕） */
@media (min-width: 1200px) {
  .responsive-input {
    width: 700px;
  }
}

/* 屏幕宽度 768px - 1199px（中等屏幕） */
@media (min-width: 768px) and (max-width: 1199px) {
  .responsive-input {
    width: 500px;
  }
}

/* 屏幕宽度 < 768px（小屏幕/手机） */
@media (max-width: 767px) {
  .responsive-input {
    width: 100%;
    /* 小屏幕时取消最大宽度限制，充分利用空间 */
    max-width: none;
  }
}

.tag {
  font-size: 16px;
  background-color: white;
}
</style>实现排个计划的功能