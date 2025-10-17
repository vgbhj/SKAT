<script setup>
const filters = {
  course: ['1', '2', '3', '4'],
  subjects: ['Мат.анализ', 'ООП', 'ЛинАл'],
  universities: ['РТУ МИРЭА', 'МИРОВОЕ', 'МИРОВОЙ ВУЗ', 'МИРНЫЙ ВУЗ']
}
</script>

<template>
  <div class="space-y-4">
    <!-- Course Filter -->
    <details class="group [&_summary::-webkit-details-marker]:hidden">
      <summary class="flex items-center justify-between gap-2 p-2 font-medium text-gray-900 bg-white rounded-lg cursor-pointer">
        <span>Курс</span>
        <span class="transition group-open:-rotate-180">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
          </svg>
        </span>
      </summary>

      <div class="mt-2 space-y-2">
        <div class="flex gap-2 flex-wrap">
          <button
            v-for="course in filters.course"
            :key="course"
            class="px-3 py-1 text-xs font-medium text-blue_main border border-blue_main rounded-full hover:bg-blue_main hover:text-white"
          >
            {{ course }}
          </button>
        </div>
      </div>
    </details>

    <!-- Subject Filter -->
    <details class="group [&_summary::-webkit-details-marker]:hidden">
      <summary class="flex items-center justify-between gap-2 p-2 font-medium text-gray-900 bg-white rounded-lg cursor-pointer">
        <span>Предмет / тема</span>
        <span class="transition group-open:-rotate-180">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
          </svg>
        </span>
      </summary>

      <div class="mt-2 space-y-2">
        <div class="flex gap-2 flex-wrap">
          <button
            v-for="subject in filters.subjects"
            :key="subject"
            class="px-3 py-1 text-xs font-medium text-blue_main border border-blue_main rounded-full hover:bg-blue_main hover:text-white"
          >
            {{ subject }}
          </button>
        </div>
      </div>
    </details>

    <!-- University Filter -->
    <details class="group [&_summary::-webkit-details-marker]:hidden">
      <summary class="flex items-center justify-between gap-2 p-2 font-medium text-gray-900 bg-white rounded-lg cursor-pointer">
        <span>ВУЗ и др. УЗ</span>
        <span class="transition group-open:-rotate-180">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
          </svg>
        </span>
      </summary>

      <div class="mt-2 space-y-2">
        <div class="flex gap-2 flex-wrap">
          <button
            v-for="uni in filters.universities"
            :key="uni"
            class="px-3 py-1 text-xs font-medium text-blue_main border border-blue_main rounded-full hover:bg-blue_main hover:text-white"
          >
            {{ uni }}
          </button>
        </div>
      </div>
    </details>
  </div>
</template>