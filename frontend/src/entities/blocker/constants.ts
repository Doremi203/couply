export enum ReportReason {
  unspecified = 'REASON_UNSPECIFIED',
  fakeProfile = 'REASON_FAKE_PROFILE',
  spam = 'REASON_SPAM',
  abuse = 'REASON_ABUSE',
  inappropriateContent = 'REASON_INAPPROPRIATE_CONTENT',
  age = 'REASON_AGE',
  other = 'REASON_OTHER',
}

export const reportOptions = {
  fakeProfile: 'Фейковый профиль',
  spam: 'Спам',
  abuse: 'Оскорбительное поведениe',
  inappropriateContent: 'Неприемлемый контент',
  age: 'Возраст',
  other: 'Другая причина',
};

export const reportToApi = {
  'Фейковый профиль': 'REASON_FAKE_PROFILE',
  Спам: 'REASON_SPAM',
  'Оскорбительное поведениe': 'REASON_ABUSE',
  'Неприемлемый контент': 'REASON_INAPPROPRIATE_CONTENT',
  Возраст: 'REASON_AGE',
  'Другая причина': 'REASON_OTHER',
};
