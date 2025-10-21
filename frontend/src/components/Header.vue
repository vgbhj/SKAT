<script setup>
import { ref, onMounted } from 'vue'
import { RouterLink, useRouter } from 'vue-router'

const isAuthenticated = ref(false)
const router = useRouter()

onMounted(async () => {
  try {
    const response = await fetch('/api/auth/check', {
      credentials: 'include'
    })
    isAuthenticated.value = response.ok
  } catch (error) {
    console.error('Auth check failed:', error)
  }
})

const handleLogout = async () => {
  try {
    await fetch('/logout', {
      method: 'POST',
      credentials: 'include'
    })
    isAuthenticated.value = false
    router.push('/login')
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

          <button
            v-if="isAuthenticated"
            type="button"
            aria-label="Add"
            class="flex items-center justify-center w-8 h-8 rounded-full bg-blue_main text-white"
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
          </button>
        </div>

        <div class="flex items-center gap-4">
          <div class="sm:flex sm:gap-4">
            <!-- Show these buttons when user is not authenticated -->
            <template v-if="!isAuthenticated">
              <RouterLink
                to="/rega"
                class="block rounded-md px-5 py-2.5 text-sm font-medium transition text-blue_main"
              >
                зарегистрироваться
              </RouterLink>

              <RouterLink
                to="/login"
                class="rounded-full bg-blue_main px-5 py-2.5 text-sm font-medium text-white sm:block"
              >
                войти
              </RouterLink>
            </template>

            <!-- Show these buttons when user is authenticated -->
            <template v-else>
              <RouterLink
                to="/profile"
                class="block rounded-md px-5 py-2.5 text-sm font-medium transition text-blue_main"
              >
                профиль
              </RouterLink>

              <button
                @click="handleLogout"
                class="rounded-full bg-blue_main px-5 py-2.5 text-sm font-medium text-white sm:block"
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