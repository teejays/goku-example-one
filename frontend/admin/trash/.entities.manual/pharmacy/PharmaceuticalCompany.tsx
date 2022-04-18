import { DefaultEntityInfo, EntityInfoRequired } from 'common/'

import { PharmacyServiceInfo } from 'types/services/'
import { StringField } from 'common/FieldKind'
import { UUID } from 'types/scalars/'

export type PharmaceuticalCompany = {
    id: UUID
    name: string
    created_at: Date
    updated_at: Date
}

const PharmaceuticalCompanyInfoLocal: EntityInfoRequired = {
    name: 'pharmaceutical_company',
    service: PharmacyServiceInfo,
    fieldInfos: [{ name: 'id', kind: StringField }],
}

export const PharmaceuticalCompanyInfo = {
    ...DefaultEntityInfo,
    ...PharmaceuticalCompanyInfoLocal,
}
