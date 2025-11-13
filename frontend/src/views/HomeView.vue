<script setup>
import { ref, onMounted } from 'vue'
import Filter from '@/components/Filter.vue'
import HomeKat from '@/components/HomeKat.vue'

const materials = ref([])
const loading = ref(false)
const error = ref('')

const fetchMaterials = async () => {
  loading.value = true
  error.value = ''
  try {
    const response = await fetch('/api/v1/materials', {
      method: 'GET',
      headers: {
        'accept': 'application/json'
      }
    })

    if (response.ok) {
      const data = await response.json()
      materials.value = data.data.materials || []
    } else {
      error.value = 'Ошибка при загрузке материалов'
    }
  } catch (err) {
    error.value = 'Ошибка сервера'
    console.error('Fetch error:', err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchMaterials()
})
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
        <h1 class="text-3xl font-bold mb-6">Последние материалы</h1>
        
        <!-- Ошибка -->
        <div v-if="error" class="mb-4 p-4 bg-red-100 text-red-700 rounded-lg">
          {{ error }}
        </div>

        <!-- Загрузка -->
        <div v-if="loading" class="text-center py-8">
          <p class="text-gray-500">Загрузка материалов...</p>
        </div>

        <!-- Материалы -->
        <div v-else-if="materials.length > 0" class="space-y-4">
          <HomeKat 
            v-for="material in materials"
            :key="material.id"
            :material="material"
          />
        </div>

        <!-- Нет материалов -->
        <div v-else class="text-center py-8">
          <p class="text-gray-500 text-lg">Материалы не найдены</p>
        </div>
      </div>
    </div>
  </main>
</template>