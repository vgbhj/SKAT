<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()

const material = ref(null)
const loading = ref(false)
const error = ref('')
const downloading = ref(false)

const fetchMaterial = async () => {
  loading.value = true
  error.value = ''
  try {
    const response = await fetch(`/api/v1/materials/${route.params.id}`, {
      method: 'GET',
      headers: {
        'accept': 'application/json'
      }
    })

    if (response.ok) {
      const data = await response.json()
      material.value = data.data.material
    } else {
      error.value = 'Ошибка при загрузке материала'
    }
  } catch (err) {
    error.value = 'Ошибка сервера'
    console.error('Fetch error:', err)
  } finally {
    loading.value = false
  }
}

const handleDownload = async () => {
  downloading.value = true
  try {
    const response = await fetch(`/api/v1/materials/${route.params.id}/download`, {
      method: 'GET'
    })

    if (response.ok) {
      const blob = await response.blob()
      const url = window.URL.createObjectURL(blob)
      const a = document.createElement('a')
      a.href = url
      a.download = material.value.filename.split('/').pop()
      document.body.appendChild(a)
      a.click()
      window.URL.revokeObjectURL(url)
      a.remove()
    } else {
      error.value = 'Ошибка при скачивании файла'
    }
  } catch (err) {
    error.value = 'Ошибка при скачивании'
    console.error('Download error:', err)
  } finally {
    downloading.value = false
  }
}

const formatDate = (dateString) => {
  const date = new Date(dateString)
  return new Intl.DateTimeFormat('ru-RU', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}

const goBack = () => {
  router.back()
}

onMounted(() => {
  fetchMaterial()
})
</script>

<template>
  <main class="mx-auto max-w-screen-xl px-4 sm:px-6 lg:px-8 py-8">
    <!-- Back button -->
    <button
      @click="goBack"
      class="mb-6 inline-flex items-center text-blue_main hover:opacity-90 transition-opacity"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="h-5 w-5 mr-2"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M15.75 19.5L8.25 12l7.5-7.5"
        />
      </svg>
      Назад
    </button>

    <!-- Loading state -->
    <div v-if="loading" class="text-center py-12">
      <p class="text-gray-500 text-lg">Загрузка материала...</p>
    </div>

    <!-- Error state -->
    <div v-else-if="error" class="bg-red-100 text-red-700 p-4 rounded-lg mb-4">
      {{ error }}
    </div>

    <!-- Material content -->
    <div v-else-if="material" class="bg-white rounded-lg shadow-lg overflow-hidden">
      <!-- Header -->
      <div class="bg-gradient-to-r from-blue_main to-blue-600 px-6 py-8 text-white">
        <h1 class="text-4xl font-bold mb-4">
          {{ material.title }}
        </h1>
        <div class="flex items-center gap-4">
          <time class="text-sm opacity-90">
            {{ formatDate(material.upload_date) }}
          </time>
        </div>
      </div>

      <!-- Content -->
      <div class="p-8">
        <!-- Description -->
        <section class="mb-8">
          <h2 class="text-2xl font-semibold text-gray-900 mb-4">Описание</h2>
          <p class="text-gray-700 text-lg leading-relaxed">
            {{ material.desc }}
          </p>
        </section>

        <!-- File info -->
        <section class="mb-8 p-4 bg-gray-50 rounded-lg">
          <h3 class="font-semibold text-gray-900 mb-2">Информация о файле</h3>
          <p class="text-gray-600 text-sm">
            <strong>Имя файла:</strong> {{ material.filename.split('/').pop() }}
          </p>
          <p class="text-gray-600 text-sm">
            <strong>Загружено:</strong> {{ formatDate(material.upload_date) }}
          </p>
        </section>

        <!-- Download button -->
        <div class="flex gap-4">
          <button
            @click="handleDownload"
            :disabled="downloading"
            class="inline-flex items-center bg-blue_main text-white px-6 py-3 rounded-lg font-semibold hover:opacity-90 transition-opacity disabled:opacity-50"
          >
            <svg
              v-if="!downloading"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="h-5 w-5 mr-2"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5M16.5 12L12 16.5m0 0L7.5 12m4.5 4.5V3"
              />
            </svg>
            <span>{{ downloading ? 'Скачивание...' : 'Скачать материал' }}</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Not found -->
    <div v-else class="text-center py-12">
      <p class="text-gray-500 text-lg">Материал не найден</p>
    </div>
  </main>
</template>