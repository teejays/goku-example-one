// Generated: Register Pharmacy Service
import { getEntityInfos } from './pharmacy'
import { registerEntityInfos } from 'common/EntityInfoStore'

// Q: When will this code be called?
export const registerEntities = () => {
    registerEntityInfos(getEntityInfos())
}
