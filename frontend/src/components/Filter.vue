<script setup>
import { ref } from 'vue'

const filters = {
  course: ['1', '2', '3', '4'],
  subjects: ['Мат.анализ', 'ООП', 'ЛинАл'],
  universities: ['РТУ МИРЭА', 'МИРОВОЕ', 'МИРОВОЙ ВУЗ', 'МИРНЫЙ ВУЗ']
}

const selectedTags = ref({
  course: ['1', '2'],
  subjects: ['Мат.анализ'],
  universities: ['РТУ МИРЭА']
})

const activeDropdown = ref(null)

const toggleDropdown = (section) => {
  activeDropdown.value = activeDropdown.value === section ? null : section
}

const closeAllDropdowns = () => {
  activeDropdown.value = null
}
</script>

 <!-- <button class="p-1 rounded-full bg-blue_main">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
          </button> -->

<template>
  <div class="overflow-hidden rounded-lg border border-gray-200 bg-white shadow-sm max-w-xs max-h-[600px] overflow-y-auto">
    <!-- Header -->
    <div class="p-4 sticky top-0 bg-white z-10">
      <h3 class="text-lg font-medium text-gray-900">Теги</h3>
      <div class="mt-2 h-px bg-gray-200"></div>
    </div>

    <!-- Filter sections -->
    <div class="px-4 pb-4 space-y-6">
      <!-- Course section -->
      <div class="space-y-2">
        <span class="text-sm font-medium text-blue_main">Курс</span>
        <div class="flex items-center gap-2">
          <div class="relative">
            <button 
              @click="toggleDropdown('course')"
              class="inline-flex items-center justify-center rounded-full bg-blue_main p-1.5 text-white hover:opacity-90 transition-opacity"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
              </svg>
            </button>

            <!-- Dropdown -->
            <div 
              v-if="activeDropdown === 'course'"
              class="fixed z-50 w-48 rounded-lg border border-gray-200 bg-white py-2 shadow-lg"
              style="margin-top: 4px;"
            >
              <div 
                v-for="option in filters.course" 
                :key="option"
                @click="toggleDropdown('course')"
                class="px-4 py-2 text-sm text-gray-600 hover:bg-[#EBF0FB] hover:text-blue_main cursor-pointer"
              >
                {{ option }}
              </div>
            </div>
          </div>
          <div class="flex flex-wrap gap-2">
            <span 
              v-for="tag in selectedTags.course" 
              :key="tag" 
              class="inline-flex h-6 w-6 items-center justify-center rounded-full  border-blue_main bg-[#EBF0FB] text-xs font-medium text-blue_main"
            >
              {{ tag }}
            </span>
          </div>
        </div>
      </div>

      <!-- Subject section -->
      <div class="space-y-2">
        <span class="text-sm font-medium text-blue_main">Предмет / тема</span>
        <div class="flex items-center gap-2">
          <div class="relative">
            <button 
              @click="toggleDropdown('subjects')"
              class="inline-flex items-center justify-center rounded-full bg-blue_main p-1.5 text-white hover:opacity-90 transition-opacity"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
              </svg>
            </button>

            <!-- Dropdown -->
            <div 
              v-if="activeDropdown === 'subjects'"
              class="fixed z-50 w-48 rounded-lg border border-gray-200 bg-white py-2 shadow-lg"
              style="margin-top: 4px;"
            >
              <div 
                v-for="option in filters.subjects" 
                :key="option"
                @click="toggleDropdown('subjects')"
                class="px-4 py-2 text-sm text-gray-600 hover:bg-[#EBF0FB] hover:text-blue_main cursor-pointer"
              >
                {{ option }}
              </div>
            </div>
          </div>
          <div class="flex flex-wrap gap-2">
            <span 
              v-for="tag in selectedTags.subjects" 
              :key="tag" 
              class="inline-flex items-center rounded-full  border-blue_main bg-[#EBF0FB] px-2 py-1 text-xs font-medium text-blue_main"
            >
              {{ tag }}
            </span>
          </div>
        </div>
      </div>

      <!-- University section -->
      <div class="space-y-2">
        <span class="text-sm font-medium text-blue_main">ВУЗ</span>
        <div class="flex items-center gap-2">
          <div class="relative">
            <button 
              @click="toggleDropdown('universities')"
              class="inline-flex items-center justify-center rounded-full bg-blue_main p-1.5 text-white hover:opacity-90 transition-opacity"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
              </svg>
            </button>

            <!-- Dropdown -->
            <div 
              v-if="activeDropdown === 'universities'"
              class="fixed z-50 w-48 rounded-lg border border-gray-200 bg-white py-2 shadow-lg"
              style="margin-top: 4px;"
            >
              <div 
                v-for="option in filters.universities" 
                :key="option"
                @click="toggleDropdown('universities')"
                class="px-4 py-2 text-sm text-gray-600 hover:bg-[#EBF0FB] hover:text-blue_main cursor-pointer"
              >
                {{ option }}
              </div>
            </div>
          </div>
          <div class="flex flex-wrap gap-2">
            <span 
              v-for="tag in selectedTags.universities" 
              :key="tag" 
              class="inline-flex items-center rounded-full  border-blue_main bg-[#EBF0FB] px-2 py-1 text-xs font-medium text-blue_main"
            >
              {{ tag }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>