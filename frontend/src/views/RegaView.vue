<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const username = ref('')  // Changed from email
const password = ref('')
const confirmPassword = ref('')
const error = ref('')

const handleSubmit = async () => {
  if (password.value !== confirmPassword.value) {
    error.value = 'Пароли не совпадают'
    return
  }

  try {
    const formData = new URLSearchParams()
    formData.append('username', username.value)
    formData.append('password', password.value)

    const response = await fetch('/signup', {
      method: 'POST',
      headers: {
        'accept': 'application/json',
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      body: formData
    })

    if (response.ok) {
      router.push('/login')
    } else {
      const data = await response.json()
      error.value = data.message || 'Ошибка регистрации'
    }
  } catch (err) {
    error.value = 'Ошибка сервера'
    console.error('Signup error:', err)
  }
}
</script>

<template>
  <main class="mx-auto max-w-screen-xl px-4 sm:px-6 lg:px-8">
    <div class="flex min-h-screen items-center justify-center">
      <div class="w-full max-w-md">
        <!-- Header -->
        <div class="mb-8 text-center">
          <h1 class="text-3xl font-bold text-gray-900">Регистрация</h1>
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

          <!-- Confirm Password field -->
          <div>
            <input
              v-model="confirmPassword"
              type="password"
              required
              placeholder="Подтвердите пароль"
              class="w-full rounded-lg border border-gray-300 px-4 py-3 text-sm focus:border-blue_main focus:outline-none focus:ring-1 focus:ring-blue_main"
            />
          </div>

          <!-- Submit button -->
          <button
            type="submit"
            class="w-full rounded-lg bg-blue_main py-3 text-sm font-semibold text-white hover:opacity-90 transition-opacity"
          >
            Зарегистрироваться
          </button>
        </form>
      </div>
    </div>
  </main>
</template>