CREATE TABLE IF NOT EXISTS uploads (
	id bigserial PRIMARY KEY,
	added_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	twitchlogin text NOT NULL,
	twitchid text NOT NULL,
	twitchmessage text NOT NULL,
	twitchchannel text NOT NULL,
	filehoster text NOT NULL,
	downloadurl text,
	uploadurl text,
	identifier text
);

INSERT INTO uploads (added_at,twitchlogin,twitchid,twitchchannel,twitchmessage,filehoster,downloadurl,uploadurl,identifier) VALUES
	 (NOW(),'nourylul','31437432','nourylul','()yaf https://www.youtube.com/watch?v=3rBFkwtaQbU','yaf','https://www.youtube.com/watch?v=3rBFkwtaQbU','https://i.yaf.ee/LEFuX.webm','a4af2284-4e13-46fa-9896-393bb1771a9d'),
	 (NOW(),'uudelleenkytkeytynyt','465178364','nourylul','()gofile https://www.youtube.com/watch?v=st6yupvNkVo','gofile','https://www.youtube.com/watch?v=st6yupvNkVo','https://gofile.io/d/PD1QNr','4ec952cc-42c0-41cd-9b07-637b4ec3c2b3');

