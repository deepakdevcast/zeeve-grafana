import { getBackendSrv } from 'app/core/services/backend_srv';
import store from 'app/core/store';
import { getDatasourceSrv } from 'app/features/plugins/datasource_srv';

import { SetupStep } from './types';

const step1TutorialTitle = 'Zeeve Observability fundamentals';
const step2TutorialTitle = 'Create users and teams';
const keyPrefix = 'getting.started.';
const step1Key = `${keyPrefix}${step1TutorialTitle.replace(' ', '-').trim().toLowerCase()}`;
const step2Key = `${keyPrefix}${step2TutorialTitle.replace(' ', '-').trim().toLowerCase()}`;

export const getSteps = (): SetupStep[] => [];
