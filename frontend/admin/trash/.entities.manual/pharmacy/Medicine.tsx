import * as Entities from './'
import * as FieldKind from 'common/FieldKind'

import { DefaultEntityInfo, EntityInfo, EntityInfoRequired } from 'common/'

import { PharmaceuticalCompanyInfo } from './PharmaceuticalCompany'
import { PharmacyServiceInfo } from 'types/services/'
import { UUID } from 'types/scalars/'

export type Medicine = {
    id: UUID
    name: string
    description: string
    active_ingredients_ids: UUID[]
    inactive_ingredients_ids: UUID[]
    company_id: UUID
    root_company_id: UUID | null // nullable
    created_at: Date
    updated_at: Date
}

const MedicineInfoLocal: EntityInfoRequired = {
    name: 'medicine',
    service: PharmacyServiceInfo,

    fieldInfos: [
        { name: 'id', kind: FieldKind.UuidField, isMeta: true },
        { name: 'name', kind: FieldKind.StringField },
        { name: 'description', kind: FieldKind.StringField },
        {
            name: 'active_ingredients_ids',
            kind: FieldKind.ForeignObjectRepeatedField,
            foreignEntityInfo: Entities.DrugInfo,
        },
        {
            name: 'inactive_ingredients_ids',
            kind: FieldKind.ForeignObjectRepeatedField,
            foreignEntityInfo: Entities.DrugInfo,
        },
        {
            name: 'company_id',
            kind: FieldKind.ForeignObjectField,
            foreignEntityInfo: PharmaceuticalCompanyInfo,
        },
        {
            name: 'root_company_id',
            kind: FieldKind.ForeignObjectField,
            foreignEntityInfo: PharmaceuticalCompanyInfo,
        },
        {
            name: 'created_at',
            kind: FieldKind.DatetimeField,
            isMeta: true,
        },
        {
            name: 'updated_at',
            kind: FieldKind.DatetimeField,
            isMeta: true,
        },
    ],
}

export const MedicineInfo: EntityInfo = {
    ...DefaultEntityInfo,
    ...MedicineInfoLocal,
}
