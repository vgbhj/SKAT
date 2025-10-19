<script setup>
import Filter from '@/components/Filter.vue';
import HomeKat from '@/components/HomeKat.vue';
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
        <h1 class="text-3xl font-bold mb-6">Последние скаты</h1>
        <div class="space-y-4"> <!-- Added container with vertical spacing -->
          <HomeKat />
          <HomeKat />
          <HomeKat />
          <HomeKat />
          <HomeKat />
        </div>
      </div>
    </div>
  </main>
</template>