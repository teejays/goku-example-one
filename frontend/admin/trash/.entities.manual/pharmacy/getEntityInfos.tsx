import { DrugInfo } from './Drug'
import { EntityInfo } from 'common'
// - Local Imports
import { MedicineInfo } from './Medicine'
import { PharmaceuticalCompanyInfo } from './PharmaceuticalCompany'

export const getEntityInfos = (): EntityInfo[] => {
    return [MedicineInfo, DrugInfo, PharmaceuticalCompanyInfo]
}
