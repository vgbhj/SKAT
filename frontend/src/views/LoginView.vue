<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const username = ref('')
const password = ref('')
const error = ref('')

const handleSubmit = async () => {
  try {
    const formData = new URLSearchParams()
    formData.append('username', username.value)
    formData.append('password', password.value)

    const response = await fetch('/login', {
      method: 'POST',
      credentials: 'include',
      headers: {
        'accept': 'application/json',
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      body: formData
    })

    if (response.ok) {
      router.push('/')
    } else {
      const data = await response.json()
      error.value = data.message || 'Ошибка входа'
    }
  } catch (err) {
    error.value = 'Ошибка сервера'
    console.error('Login error:', err)
  }
}
</script>

<template>
  <main class="mx-auto max-w-screen-xl px-4 sm:px-6 lg:px-8">
    <div class="flex min-h-screen items-center justify-center">
      <div class="w-full max-w-md">
        <!-- Header -->
        <div class="mb-8 text-center">
          <h1 class="text-3xl font-bold text-gray-900">Вход</h1>
        </div>

        <!-- Error message -->
        <div v-if="error" class="mb-4 text-red-500 text-center">
          {{ error }}
        </div>

        <!-- Form -->
        <form @submit.prevent="handleSubmit" class="space-y-6">
          <!-- Username field -->
          <div>
            <input
              v-model="username"
              type="text"
              required
              placeholder="Введите имя пользователя"
              class="w-full rounded-lg border border-gray-300 px-4 py-3 text-sm focus:border-blue_main focus:outline-none focus:ring-1 focus:ring-blue_main"
            />
          </div>

          <!-- Password field -->
          <div>
            <input
              v-model="password"
              type="password"
              required
              placeholder="Введите пароль"
              class="w-full rounded-lg border border-gray-300 px-4 py-3 text-sm focus:border-blue_main focus:outline-none focus:ring-1 focus:ring-blue_main"
            />
          </div>

          <!-- Submit button -->
          <button
            type="submit"
            class="w-full rounded-lg bg-blue_main py-3 text-sm font-semibold text-white hover:opacity-90 transition-opacity"
          >
            Войти
          </button>
        </form>
      </div>
    </div>
  </main>
</template>