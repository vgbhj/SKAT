<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const title = ref('')
const description = ref('')
const file = ref(null)
const error = ref('')
const loading = ref(false)

const handleFileChange = (e) => {
  file.value = e.target.files[0]
}

const handleSubmit = async () => {
  if (!file.value) {
    error.value = 'Пожалуйста, выберите файл'
    return
  }

  loading.value = true
  error.value = ''

  try {
    const formData = new FormData()
    formData.append('title', title.value)
    formData.append('desc', description.value)
    formData.append('file', file.value)

    const response = await fetch('/api/v1/materials', {
      method: 'POST',
      credentials: 'include',
      headers: {
        'accept': 'application/json'
      },
      body: formData
    })

    if (response.ok) {
      router.push('/')
    } else {
      const data = await response.json()
      error.value = data.message || 'Ошибка при загрузке материала'
    }
  } catch (err) {
    error.value = 'Ошибка сервера'
    console.error('Upload error:', err)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <main class="mx-auto max-w-screen-xl px-4 sm:px-6 lg:px-8">
    <div class="flex min-h-screen items-center justify-center">
      <div class="w-full max-w-md">
        <!-- Header -->
        <div class="mb-8 text-center">
          <h1 class="text-3xl font-bold text-gray-900">Добавить материал</h1>
        </div>

        <!-- Error message -->
        <div v-if="error" class="mb-4 text-red-500 text-center">
          {{ error }}
        </div>

        <!-- Form -->
        <form @submit.prevent="handleSubmit" class="space-y-6">
          <!-- Title field -->
          <div>
            <input
              v-model="title"
              type="text"
              required
              placeholder="Название материала"
              class="w-full rounded-lg border border-gray-300 px-4 py-3 text-sm focus:border-blue_main focus:outline-none focus:ring-1 focus:ring-blue_main"
            />
          </div>

          <!-- Description field -->
          <div>
            <textarea
              v-model="description"
              required
              placeholder="Описание материала"
              rows="4"
              class="w-full rounded-lg border border-gray-300 px-4 py-3 text-sm focus:border-blue_main focus:outline-none focus:ring-1 focus:ring-blue_main"
            ></textarea>
          </div>

          <!-- File upload -->
          <div>
            <input
              type="file"
              @change="handleFileChange"
              required
              class="w-full rounded-lg border border-gray-300 px-4 py-3 text-sm focus:border-blue_main focus:outline-none focus:ring-1 focus:ring-blue_main"
            />
          </div>

          <!-- Submit button -->
          <button
            type="submit"
            :disabled="loading"
            class="w-full rounded-lg bg-blue_main py-3 text-sm font-semibold text-white hover:opacity-90 transition-opacity disabled:opacity-50"
          >
            {{ loading ? 'Загрузка...' : 'Добавить материал' }}
          </button>
        </form>
      </div>
    </div>
  </main>
</template>