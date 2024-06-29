<script setup lang="ts">
    import type { UserRepository as UserRepo } from '../../domain/contract/userContract'
    import { useUser } from '../../composable/userComposable'

    const props = defineProps<{
		userRepositoy: UserRepo,
        userId?: string,
	}>();

    const { 
        userForm,
		saveUser,
        getUserById,
        setUserData,
	} = useUser({
        userRepository: props.userRepositoy,
        userId: props?.userId
    });	

    const loadInitialData = () => {

        if(props?.userId) {
            getUserById(props.userId).then((user) => {
                if(user) {
                    setUserData(user);
                }
            });
        }
    }
    
    loadInitialData();
    defineExpose({ save: saveUser });
</script>

<template>
    <form ref="formRef">
    <div class="flexbox-grid container">
        <div class="card" style="--min: 50ch">
            <div class="grid">
                <h2>{{  $t('SHARED.basic_information') }}</h2>

				<span class="p-float-label">
					<InputText v-model="userForm.name" />
                    <label for="name">{{ $t('SECURITY.name') }}</label>
                </span>

                <span class="p-float-label">
					<InputText v-model="userForm.lastName" />
                    <label for="last_name">{{ $t('SECURITY.last_name') }}</label>
                </span>
			</div>
        </div>

        <div class="card" style="--min: 35ch">
            <div class="grid">
                <h2>{{ $t('SECURITY.security') }}</h2>

				<span class="p-float-label">
					<InputText v-model="userForm.userName" :disabled="userForm.id != '' && userForm.id != undefined"/>
                    <label for="user_name">{{ $t('SECURITY.user_name') }}</label>
                </span>

                <span class="p-float-label">
					<InputText type="password" v-model="userForm.password" :disabled="userForm.id != '' && userForm.id != undefined" />
                    <label for="password">{{ $t('SECURITY.password') }}</label>
                </span>

                <span class="p-float-label">
					<InputText type="password" v-model="userForm.confirmationPassword" :disabled="userForm.id != '' && userForm.id != undefined" />
                    <label for="password_confirmation">{{ $t('SECURITY.password_confirmation') }}</label>
                </span>
			</div>
        </div>
    </div>

    </form>
</template>

<style scoped>

</style>