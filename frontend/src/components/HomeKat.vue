<script setup>
import { useRouter } from 'vue-router'

defineProps({
  material: {
    type: Object,
    required: true
  }
})

const router = useRouter()

const post = {
  user: {
    avatar: 'https://ui-avatars.com/api/?name=User+Name&background=0D8ABC&color=fff',
    name: 'User Name'
  },
  tags: ['Мат.анализ', 'РТУ МИРЭА', '1 курс']
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

const goToMaterial = () => {
  router.push({ name: 'material', params: { id: material.id } })
}
</script>

<template>
    <article 
    @click="goToMaterial"
    class="rounded-lg border border-gray-200 bg-white p-4 shadow-sm hover:shadow-lg hover:cursor-pointer transition-all duration-200"
  >
    <!-- Header with avatar and metadata -->
    <header class="mb-4">
      <div class="flex items-center gap-4">
        <!-- Avatar -->
        <img 
          :src="post.user.avatar" 
          :alt="post.user.name"
          class="h-6 w-6 rounded-full object-cover"
        />
        
        <div class="flex-1">
          <div class="flex items-center justify-between">
            <!-- Tags -->
            <div class="flex flex-wrap gap-2">
              <span 
                v-for="tag in post.tags" 
                :key="tag"
                class="inline-flex items-center rounded-full border-blue_main bg-[#EBF0FB] px-2 py-1 text-xs font-medium text-blue_main"
              >
                {{ tag }}
              </span>
            </div>
            <!-- Timestamp -->
            <time class="text-sm text-gray-500">
              {{ formatDate(material.upload_date) }}
            </time>
          </div>
        </div>
      </div>
    </header>

    <!-- Description -->
    <div class="pl-10 cursor-pointer">
      <!-- Title -->
      <h3 class="text-lg font-semibold text-gray-900 mb-2 hover:text-blue_main transition-colors">
        {{ material.title }}
      </h3>
      <!-- Description -->
      <p class="text-gray-700 font-medium">
        {{ material.desc }}
      </p>
    </div>
  </article>
</template>