<script setup>
import { ref } from 'vue'

const email = ref('')
const password = ref('')

const handleSubmit = () => {
  // Add your login logic here
  console.log('Form submitted:', {
    email: email.value,
    password: password.value
  })
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

        <!-- Form -->
        <form @submit.prevent="handleSubmit" class="space-y-6">
          <!-- Email field -->
          <div>
            <input
              v-model="email"
              type="email"
              required
              placeholder="Введите e-mail"
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