<script setup>
import Filter from '@/components/Filter.vue';
</script>

<template>
  <main class="mx-auto max-w-screen-xl px-4 sm:px-6 lg:px-8">
    <div class="grid grid-cols-12 gap-8">
      <!-- Левая колонка с фильтрами -->
      <div class="col-span-3">
        <Filter />
      </div>
      
      <!-- Правая колонка с контентом -->
      <div class="col-span-9">
        <!-- Здесь будет основной контент -->
      </div>
    </div>
  </main>
</template>