Onboarding 1 - Full
1. Select languages (multiple, only complete lang's)
2. Select accent (one each lang)

Onboarding 2 - With surah selection
1. Select surahs
2. Select language (multiple, can be partial if covers the surah)
3. Select accent (one each lang)

Tables & relations
1. Surahs
    1. Al-Fatiha
    2. ...
2. Ayats
    1.  #1
    2. ...
3. Language
    1. Arabic
    2. ...
4. Accent
    1. #1
    2. ...
5. Audio - ?

# Goal: get 1st ayat of first surah that is in arabic language and accent 1
Data we've:
1. ayat_id
2. lng_id
3. accent_id

select * from "audios"
left join "lng"
on
where
ayat_id = 1
accent_id = 1
    