import { set as setStore, get as gesStore } from '@/modules/shared/utility/localStorage'
import { ref } from 'vue'

export const useLanguage = (globalProperties: any) => {
	const currentLang = ref(gesStore('lang') || 'en');
	
    const changeLanguage = (language: 'es' | 'en') => {
		const languageChage = globalProperties.changeLanguage;
		if(!languageChage) return;

		currentLang.value = language;

		
		languageChage(language);
		setStore('lang', language);
    }
	
    return {
		currentLang,
        changeLanguage
    }
}
