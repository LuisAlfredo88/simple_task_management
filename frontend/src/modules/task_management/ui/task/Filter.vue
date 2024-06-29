<script setup lang="ts">
    import type { TaskFilter } from "../../domain/entity/task"
    import Dropdown from 'primevue/dropdown'
    import { ref } from 'vue'
    
    const emit = defineEmits(['update:modelValue', 'onChage']);
    
    const props = defineProps<{
		modelValue: TaskFilter,
        taskStatus: GenericKeyValue[]
	}>();

    const filter = ref<TaskFilter>({...props.modelValue});

    const clearFilter = () => {
        filter.value = {} as TaskFilter;
        emit('update:modelValue', filter);
        emit('onChage', filter);
    }
    
    const appy = () => {
        emit('update:modelValue', filter);
        emit('onChage', filter);
    }
</script>

<template>
    <div class="filter-container">
        <div class="flexbox-grid">
            <span class="p-float-label">
                <Dropdown placeholder="Select an item" :showClear="true" optionLabel="name" v-model="filter.statusId" optionValue="id" :options="taskStatus" filter  />
                <label for="Origin">{{ $t('SECURITY.status') }}</label>
            </span>
        </div>

        <div class="filter-actions">
            <PButton :label="$t('COMMON_BUTTONS.clean')" @click="clearFilter" severity="warning" icon="pi pi-eraser" autofocus />
            <PButton :label="$t('COMMON_BUTTONS.apply')" @click="appy" icon="pi pi-check" autofocus />
        </div>
    </div>
</template>