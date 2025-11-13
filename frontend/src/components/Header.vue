<script setup>
import { useAuthStore } from '@/stores/auth'
import { useRouter } from 'vue-router'

const authStore = useAuthStore()
const router = useRouter()

const handleLogout = async () => {
  try {
    const response = await fetch('/logout', {
      method: 'POST',
      credentials: 'include'
    })
    if (response.ok) {
      authStore.logout()
      router.push('/login')
    }
  } catch (error) {
    console.error('Logout failed:', error)
  }
}
</script>

<template>
  <header class="">
    <div class="mx-auto flex h-16 max-w-screen-xl items-center gap-8 px-4 sm:px-6 lg:px-8">
      <RouterLink class="block text-teal-600 dark:text-teal-300" to="/">
        <span class="inline-flex items-center justify-center rounded-full bg-blue_main px-2.5 py-0.5 text-white">
          <p class="whitespace-nowrap p-1 font-medium">SKAT</p>
        </span>
      </RouterLink>

      <div class="flex flex-1 items-center justify-end md:justify-between">
        <div class="flex items-center space-x-2">
          <label for="Search">
            <div class="relative">
              <input
                type="text"
                id="Search"
                class="w-full rounded-full border-2 border-blue_main p-1"
                placeholder="Поиск..."
              />
              <span class="absolute inset-y-0 right-1 grid w-8 place-content-center">
                <button
                  type="button"
                  aria-label="Submit"
                  class="rounded-full p-1.5 text-white transition-colors bg-blue_main"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    class="size-4"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z"
                    />
                  </svg>
                </button>
              </span>
            </div>
          </label>

          <!-- Кнопка добавления материала (видна только для аутентифицированных) -->
          <RouterLink
            v-if="authStore.isAuthenticated"
            to="/materials/add"
            type="button"
            aria-label="Add"
            class="flex items-center justify-center w-8 h-8 rounded-full bg-blue_main text-white hover:opacity-90 transition-opacity"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="h-6 w-6"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 4v16m8-8H4"
              />
            </svg>
          </RouterLink>
        </div>

        <div class="flex items-center gap-4">
          <div class="sm:flex sm:gap-4">
            <!-- Кнопки для неаутентифицированных пользователей -->
            <template v-if="!authStore.isAuthenticated">
              <RouterLink
                to="/rega"
                class="block rounded-md px-5 py-2.5 text-sm font-medium transition text-blue_main hover:bg-gray-100"
              >
                зарегистрироваться
              </RouterLink>

              <RouterLink
                to="/login"
                class="rounded-full bg-blue_main px-5 py-2.5 text-sm font-medium text-white sm:block hover:opacity-90 transition-opacity"
              >
                войти
              </RouterLink>
            </template>

            <!-- Кнопки для аутентифицированных пользователей -->
            <template v-else>
              <div class="flex items-center gap-2 px-4 py-2.5 text-sm font-medium text-gray-700 border-r border-gray-200">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke-width="1.5"
                  stroke="currentColor"
                  class="h-5 w-5"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="M15.75 6a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0zM4.501 20.118a7.5 7.5 0 0114.998 0A17.933 17.933 0 0112 21.75c-2.676 0-5.216-.584-7.499-1.632z"
                  />
                </svg>
                <span>{{ authStore.user?.username || 'Пользователь' }}</span>
              </div>

              <RouterLink
                to="/profile"
                class="block rounded-md px-5 py-2.5 text-sm font-medium transition text-blue_main hover:bg-gray-100"
              >
                профиль
              </RouterLink>

              <button
                @click="handleLogout"
                class="rounded-full bg-blue_main px-5 py-2.5 text-sm font-medium text-white sm:block hover:opacity-90 transition-opacity"
              >
                выйти
              </button>
            </template>
          </div>
        </div>
      </div>
    </div>
  </header>
</template>