import { DefaultEntityInfo, EntityInfoRequired } from 'common/'

import { PharmacyServiceInfo } from 'types/services/'
import { StringField } from 'common/FieldKind'
import { UUID } from 'types/scalars/'

export type Drug = {
    id: UUID
    name: string
    created_at: Date
    updated_at: Date
}

const DrugInfoLocal: EntityInfoRequired = {
    name: 'drug',
    service: PharmacyServiceInfo,
    fieldInfos: [{ name: 'id', kind: StringField }],
}

export const DrugInfo = { ...DefaultEntityInfo, ...DrugInfoLocal }
