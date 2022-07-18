-- Historical and upcoming commission schedule changes

BEGIN;

-- Commissions Amendments

CREATE TABLE IF NOT EXISTS oasis_3.commissions
(
  address  TEXT PRIMARY KEY NOT NULL REFERENCES oasis_3.accounts(address),
  schedule JSON
);

TRUNCATE oasis_3.commissions CASCADE;
INSERT INTO oasis_3.commissions (address, schedule) VALUES
	('oasis1qq4f2h225gv6g8w8w23fm740aze9lke4qun72n59', '{"rates":[{"start":354,"rate":"13000"},{"start":1510,"rate":"17000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzjp96xfmfauvpe6sz99r06q9tah8utlc5lkqr9w', '{"rates":[{"start":11650,"rate":"5000"},{"start":11744,"rate":"2500"}],"bounds":[{"start":11650,"rate_min":"0","rate_max":"7000"}]}'),
	('oasis1qqrv4g5wu543wa7fcae76eucqfn2uc77zgqw8fxk', '{"rates":[{"start":9438,"rate":"19000"},{"start":10188,"rate":"20000"}],"bounds":[{"start":9438,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qp8sn8t0fk727n2lwr8zzxwu5gyf7jy9xus6p4aq', '{"rates":[{"start":1599,"rate":"1000"}],"bounds":[{"start":1599,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzthup6qts0k689z2wy84yvk9ctnht66eyxl7268', '{"rates":[{"rate":"5000"},{"start":2855,"rate":"10000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":2855,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qz0pvg26eudajp60835wl3jxhdxqz03q5qt9us34', '{"rates":[{"start":13323,"rate":"5000"}],"bounds":[{"start":13323,"rate_min":"0","rate_max":"5000"}]}'),
	('oasis1qrxyndkhehffdme39urcp2v7m2t7k06xwsuyaxqq', '{"rates":[{"rate":"5000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qrugz89g5esmhs0ezer0plsfvmcgctge35n32vmr', '{"rates":[{"start":8400,"rate":"5000"}],"bounds":[{"start":8400,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qqpxy2vp0u2ym4jltdafgrjxuh9y9fap9utzr2wg', '{"rates":[{"start":11232,"rate":"20000"}],"bounds":[{"start":11232,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qp0xuvw2a93w4yp8jwthfz93gxy87u7hes9eu2ev', '{"rates":[{"start":7760,"rate":"5000"},{"start":10486,"rate":"15000"}],"bounds":[{"start":7760,"rate_min":"0","rate_max":"5000"},{"start":10486,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qqzh32kr72v7x55cjnjp2me0pdn579u6as38kacz', '{"rates":[{"start":10767,"rate":"5000"},{"start":10923,"rate":"5000"}],"bounds":[{"start":10767,"rate_min":"5000","rate_max":"5000"},{"start":10923,"rate_min":"0","rate_max":"10000"}]}'),
	('oasis1qzzryyckptmgxxnyvt05twjufhl3ah0qtgcf4n8l', '{"rates":[{"start":8452,"rate":"10000"}],"bounds":[{"start":8452,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qryc94hn6hucev6ex79ceheve2pjesenc50svvvp', '{"rates":[{"rate":"5000"},{"start":3150,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qp4f47plgld98n5g2ltalalnndnzz96euv9n89lz', '{"rates":[{"start":419,"rate":"19500"},{"start":2026,"rate":"20000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qqxg9s39jfqwhwnnv3nlpfzvdv9hqw5n2gdvuj3r', '{"rates":[{"rate":"5000"},{"start":3051,"rate":"20000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":3051,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qr3w66akc8ud9a4zsgyjw2muvcfjgfszn5ycgc0a', '{"rates":[{"start":2310,"rate":"5000"},{"start":4371,"rate":"20000"}],"bounds":[{"start":2310,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qqxqhx9t52rsevhhtfspdxp4gsaft6ewyyeqnqy3', '{"rates":[{"start":7006,"rate":"20000"}],"bounds":[{"start":7006,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qr9zuf3n8g3znm786st3sldfw677pk3a6v85w9ds', '{"rates":[{"start":9479,"rate":"19000"},{"start":10233,"rate":"20000"}],"bounds":[{"start":9707,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qq3fq8hxrlq6pedw0q3f57daea43a6v7q5rwf0ll', '{"rates":[{"start":224,"rate":"14000"},{"start":659,"rate":"19500"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzv2rktycpcykhudyvg5l6u75v08lpfcw5nt7aj5', '{"rates":[{"start":1860,"rate":"12000"}],"bounds":[{"start":1860,"rate_min":"0","rate_max":"100000"}]}'),
	('oasis1qr87t7j3csez6jknwr4ksjn4u2pwye2ku5xjcjcc', '{"rates":[{"rate":"5000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qpn83e8hm3gdhvpfv66xj3qsetkj3ulmkugmmxn3', '{"rates":[{"rate":"5000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qpqz8g88kvw49m402k8m2r6nv4p62vsdkv5d0u6r', '{"rates":[{"start":8201,"rate":"15000"},{"start":9419,"rate":"20000"}],"bounds":[{"start":7805,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qp4ajz4vdgx3ze42ulk7m0jkxfstqcl87qymg9nx', '{"rates":[{"start":5600,"rate":"10000"}],"bounds":[{"start":5600,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzp84num6xgspdst65yv7yqegln6ndcxmuuq8s9w', '{"rates":[{"start":1171,"rate":"10000"},{"start":8148,"rate":"20000"}],"bounds":[{"start":527,"rate_min":"0","rate_max":"100000"}]}'),
	('oasis1qqzvrnpu4kw69wedg5g7mf8jy5tzuhu9vchpyh0j', '{"rates":[{"start":1599,"rate":"2000"}],"bounds":[{"start":1599,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qrw82ag2sypeytse9x9k4uxym53l5lc5jyfs2sxv', '{"rates":[{"start":10455,"rate":"20610"},{"start":11468,"rate":"19610"},{"start":11481,"rate":"19610"}],"bounds":[{"start":10455,"rate_min":"0","rate_max":"50000"},{"start":11468,"rate_min":"0","rate_max":"50000"},{"start":11481,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qpsnzv8qz4fu3lwps2tc3eg5pnryzl4h7cqxruzf', '{"rates":[{"start":89,"rate":"15000"},{"start":6735,"rate":"20000"}],"bounds":[{"start":425,"rate_min":"0","rate_max":"20000"},{"start":6735,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qps7mh7usg7u4t35ujr0l8dxjs2ly2swhu9v0mr0', '{"rates":[{"rate":"5000"},{"start":4353,"rate":"20000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qqekv2ymgzmd8j2s2u7g0hhc7e77e654kvwqtjwm', '{"rates":[{"start":4725,"rate":"10000"}],"bounds":[{"start":4725,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qqa3d6e2h3dcdj38rnpgk04grcf6p4weh534vmfs', '{"rates":[{"rate":"5000"},{"start":2114,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":2114,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qps9drw07z0gmh5z2pn7zwl3z53ate2yvqf3uzq5', '{"rates":[{"start":8103,"rate":"17000"}],"bounds":[{"start":8103,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qresj0vhmwawll6fe2vw2nlapkp6nj6etcx7a32h', '{"rates":[{"rate":"5000"},{"start":1260,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":1596,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qzpltjryl38jdmd03309z8z87gcl32v53q6xshwq', '{"rates":[{"start":3156,"rate":"15000"}],"bounds":[{"start":3156,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qqx820g2geqzeyeyfnm5hgz72eaj9emajgqmscy0', '{"rates":[{"start":26,"rate":"10000"},{"start":418,"rate":"18000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qp2vdltxqvglrwj6qur8s4dw97485vm6actrkcsm', '{"rates":[{"start":12410,"rate":"14900"}],"bounds":[{"start":12410,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qqm6hnxzwm5h04zqh4fm0w2eygx0kuj9rymg48m5', '{"rates":[{"start":7607,"rate":"7000"}],"bounds":[{"start":7607,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qppjm5sxqwps4dpyekdvz530sjmq3e5eusp7hdan', '{"rates":[{"start":961,"rate":"15000"},{"start":1706,"rate":"19500"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qpujya6e0m08y8fmvnpmvxur3htaves5sgwumzgg', '{"rates":[{"start":3816,"rate":"3000"}],"bounds":[{"start":3816,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qqnmppt4j5d2yl584euhn6g2cw9gewdswga9frg4', '{"rates":[{"start":283,"rate":"15000"},{"start":3857,"rate":"18000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qp4rp7adhegfktyg4aq3w6jelqumx6klfv5t7kvv', '{"rates":[{"rate":"5000"},{"start":30,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":340,"rate_min":"0","rate_max":"30000"}]}'),
	('oasis1qzgq240adkyxhm6598ex2jvgs7lyx5wmtg7cszmk', '{"rates":[{"rate":"5000"},{"start":1259,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":1595,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qr8al5vcpqzjspdl8yt27fqc3pydz4alhs0xqp5e', '{"rates":[{"start":5217,"rate":"20000"}],"bounds":[{"start":5217,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzp8e4ldf9zq27jle847nawll5jwwa9x75y6spaz', '{"rates":[{"rate":"5000"},{"start":1705,"rate":"10000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qp60saapdcrhe5zp3c3zk52r4dcfkr2uyuc5qjxp', '{"rates":[{"start":3,"rate":"1000"},{"start":1172,"rate":"10000"},{"start":1173,"rate":"3000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qrmexg6kh67xvnp7k42sx482nja5760stcrcdkhm', '{"rates":[{"start":26,"rate":"10000"},{"start":418,"rate":"17000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qpg5vq3jt6djfw5rq9l8kcwkt4s7vmjn4vvxujwv', '{"rates":[{"start":12119,"rate":"0"}],"bounds":[{"start":12119,"rate_min":"0","rate_max":"5000"}]}'),
	('oasis1qz8vfnkcc48grazt83gstfm6yjwyptalny8cywtp', '{"rates":[{"start":10840,"rate":"20000"},{"start":13320,"rate":"20000"}],"bounds":[{"start":10840,"rate_min":"0","rate_max":"25000"},{"start":13320,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qq3xrq0urs8qcffhvmhfhz4p0mu7ewc8rscnlwxe', '{"rates":[{"rate":"5000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzl58e7v7pk50h66s2tv3u9rzf87twp7pcv7hul6', '{"rates":[{"start":3175,"rate":"9999"},{"start":12259,"rate":"5000"}],"bounds":[{"start":3175,"rate_min":"999","rate_max":"100000"}]}'),
	('oasis1qrgxl0ylc7lvkj0akv6s32rj4k98nr0f7smf6m4k', '{"rates":[{"start":2258,"rate":"18000"},{"start":4420,"rate":"20000"}],"bounds":[{"start":2878,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzu358mpd4z5frmrq6vnwq87cqfvdmfxh5ax57cj', '{"rates":[{"start":6367,"rate":"3000"},{"start":11136,"rate":"5000"}],"bounds":[{"start":1585,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qq0xmq7r0z9sdv02t5j9zs7en3n6574gtg8v9fyt', '{"rates":[{"start":2275,"rate":"20000"},{"start":11350,"rate":"2000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":11685,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qr0jwz65c29l044a204e3cllvumdg8cmsgt2k3ql', '{"rates":[{"rate":"5000"},{"start":82,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":353,"rate_min":"0","rate_max":"20000"},{"start":420,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzt4fvcc6cw9af69tek9p3mfjwn3a5e5vcyrw7ac', '{"rates":[{"start":11231,"rate":"18000"}],"bounds":[{"start":11231,"rate_min":"0","rate_max":"18000"}]}'),
	('oasis1qzugextrcdueshq63w7l9x4xglnusznsgqa95w7e', '{"rates":[{"start":418,"rate":"20000"},{"start":454,"rate":"19000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qr5f26k9rnfa2pg6wgsuljyq7lecej3gaqqhyra5', '{"rates":[{"rate":"5000"},{"start":1310,"rate":"18000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qq5huvjkenm46zzsvgdt6evgjq6h7xp92cdh6826', '{"rates":[{"start":5148,"rate":"10000"},{"start":5335,"rate":"500"}],"bounds":[{"start":5148,"rate_min":"0","rate_max":"25000"},{"start":5671,"rate_min":"0","rate_max":"7500"}]}'),
	('oasis1qppctxzn8djkqfvrxugak9v7dp25vddq7sxqhkry', '{"rates":[{"start":8103,"rate":"15000"},{"start":10805,"rate":"19000"}],"bounds":[{"start":8103,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qrt07xfajnree25q27dxnwvmxs7drj4g3sknwmh4', '{"rates":[{"start":7544,"rate":"10000"}],"bounds":[{"start":7544,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qrflth3g7k0ymkut2zrca3ktagw6g882yvqmgzdv', '{"rates":[{"start":11182,"rate":"10000"}],"bounds":[{"start":11182,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qqyynj90zkvyhja33w4ltgej45pr48f45ymmnsrx', '{"rates":[{"rate":"5000"},{"start":1167,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":1503,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qpxpnxxk4qcgl7n55tx0yuqmrcw5cy2u5vzjq5u4', '{"rates":[{"start":5886,"rate":"5000"},{"start":6210,"rate":"19000"}],"bounds":[{"start":5113,"rate_min":"5000","rate_max":"20000"}]}'),
	('oasis1qram2p9w3yxm4px5nth8n7ugggk5rr6ay5d284at', '{"rates":[{"start":2300,"rate":"15000"},{"start":3400,"rate":"19000"}],"bounds":[{"start":2300,"rate_min":"0","rate_max":"20000"},{"start":3400,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzz8zmehzgcynlueydeqfax9cznzdw3lvgark5h3', '{"rates":[{"rate":"5000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qp6ela294e5nkay37n039fw8d50k9kygesxa8mzt', '{"rates":[{"rate":"5000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzaracvtpgp87frlk9jshdcu8aczpsrz8y94jjg4', '{"rates":[{"rate":"5000"},{"start":1260,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":1596,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qrzkqzvlw9xgyxhhqps746ssyhn5lkqrmcgz8amq', '{"rates":[{"start":1167,"rate":"10000"},{"start":1283,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qp9xlxurlcx3k5h3pkays56mp48zfv9nmcf982kn', '{"rates":[{"start":10243,"rate":"20000"}],"bounds":[{"start":10242,"rate_min":"0","rate_max":"25000"},{"start":10603,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzzytegg6jc7hxu6y8feuzkgmr75ms7hc54mz85p', '{"rates":[{"start":7734,"rate":"10000"},{"start":10542,"rate":"20000"}],"bounds":[{"start":7734,"rate_min":"0","rate_max":"25000"},{"start":10542,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qz0ea28d8p4xk8xztems60wq22f9pm2yyyd82tmt', '{"rates":[{"rate":"5000"},{"start":1195,"rate":"10000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qptk3ydjxuq3eenqwljdy45uxdye5tfg4sqar0fz', '{"rates":[{"rate":"5000"},{"start":1259,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":1595,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qz6j6elhypc70gv8faax3rlpv8ygx39grc55lwwm', '{"rates":[{"rate":"5000"},{"start":4371,"rate":"20000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qz86vltcdhjurzuvzfhkku4yaf7vf2umdvpwmtlv', '{"rates":[{"rate":"5000"},{"start":2047,"rate":"10000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qqm6l2msa0lhp080wmd0xrudzpusqjp4puuusup5', '{"rates":[{"rate":"5000"},{"start":4371,"rate":"20000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzufmg23p7jprqmptvz0vyn69c7lk6vfwuz8xapf', '{"rates":[{"rate":"5000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzf7u0w4ueuapvppslgufy02hznef0l8lvhp4q9p', '{"rates":[{"start":10784,"rate":"5000"}],"bounds":[{"start":10784,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qqw05utlqvf2ska0fyjf5yr7peg2z4tuxcjmqztp', '{"rates":[{"start":9478,"rate":"19000"},{"start":10176,"rate":"20000"}],"bounds":[{"start":9533,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzf03q57jdgdwp2w7y6a8yww6mak9khuag9qt0kd', '{"rates":[{"start":10515,"rate":"10000"},{"start":11335,"rate":"15000"}],"bounds":[{"start":10515,"rate_min":"0","rate_max":"15000"}]}'),
	('oasis1qpfcsun7zju6ku7d2mdh54j9nsmxvj76uqk35w57', '{"rates":[{"start":11074,"rate":"15000"}],"bounds":[{"start":11074,"rate_min":"1000","rate_max":"19000"}]}'),
	('oasis1qz22xm9vyg0uqxncc667m4j4p5mrsj455c743lfn', '{"rates":[{"start":45,"rate":"15000"},{"start":2031,"rate":"19000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qq4jqh66ga62pe9td5zsnfge3c9rfp6zucjr03q8', '{"rates":[{"start":10728,"rate":"10000"}],"bounds":[{"start":10728,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qq2kqzr4q942x44st97n66nmmyh7dhsuvsqyc22u', '{"rates":[{"start":224,"rate":"14000"},{"start":660,"rate":"19000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qqs5wnxvsk009swtt7ehm5fslxve96kczszwt47s', '{"rates":[{"start":2303,"rate":"20000"}],"bounds":[{"start":801,"rate_min":"0","rate_max":"100000"},{"start":2878,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qpwrq93z8s9ytu2hfjtqggc9edgwfadzevs3trvm', '{"rates":[{"start":10217,"rate":"17000"},{"start":10223,"rate":"20000"}],"bounds":[{"start":9372,"rate_min":"0","rate_max":"20000"},{"start":10552,"rate_min":"0","rate_max":"20000"},{"start":10558,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qp3fd5rplv5uh6nc9vn85ejeaszewjv2xgvfts07', '{"rates":[{"rate":"5000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qqr8y5cez0aczdlnfp9fre82whjsdgqgd5xxtv6p', '{"rates":[{"start":8103,"rate":"15000"},{"start":10804,"rate":"19000"}],"bounds":[{"start":8103,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qrc8s2trrm9zgha8wq636yetx7sxjf7x35pf3vrc', '{"rates":[{"start":5932,"rate":"5000"}],"bounds":[{"start":5932,"rate_min":"0","rate_max":"15000"}]}'),
	('oasis1qrdx0n7lgheek24t24vejdks9uqmfldtmgdv7jzz', '{"rates":[{"start":22,"rate":"15000"},{"start":2402,"rate":"20000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qztpm422n7v98f4s7ja0wt5jey7fy3xpg5ye2vtl', '{"rates":[{"rate":"5000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qp334gzlzrap6k2ch6wc9vxxplw9sg3v9cfvvgsy', '{"rates":[{"start":11255,"rate":"17000"}],"bounds":[{"start":11255,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzv7a0gkxwpfelv985fvkl24k7jh3arfwy84zw7q', '{"rates":[{"start":355,"rate":"13000"},{"start":1510,"rate":"17000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qz8u0zqhcmgxalnjgwj6m0sg6wuy2vrsgyna2z7t', '{"rates":[{"rate":"5000"},{"start":2525,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qpl883gp995zs9n6a279tqsavnaxxf0rzcdlauwu', '{"rates":[{"start":8075,"rate":"5000"}],"bounds":[{"start":8075,"rate_min":"0","rate_max":"10000"}]}'),
	('oasis1qq58f3mqxt6htvtxcayt4zfshysj36zksvwkmjg9', '{"rates":[{"start":9305,"rate":"10000"}],"bounds":[{"start":9305,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qzsp62l07fqsxgdeqszwz8hm34hhwem9ny73qnpr', '{"rates":[{"rate":"5000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzl99wft8jtt7ppprk7ce7s079z3r3t77s6pf3dd', '{"rates":[{"start":50,"rate":"15000"},{"start":2275,"rate":"20000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qqf6wmc0ax3mykd028ltgtqr49h3qffcm50gwag3', '{"rates":[{"rate":"5000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzc687uuywnel4eqtdn6x3t9hkdvf6sf2gtv4ye9', '{"rates":[{"start":40,"rate":"15000"},{"start":3719,"rate":"19000"}],"bounds":[{"start":340,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qpaygvzwd5ffh2f5p4qdqylymgqcvl7sp5gxyrl3', '{"rates":[{"start":10448,"rate":"18000"},{"start":11650,"rate":"19000"}],"bounds":[{"start":10448,"rate_min":"0","rate_max":"20000"},{"start":11650,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qr2zw0l903skt999tryu2e99v7t4g20hxc5nka9s', '{"rates":[{"start":2853,"rate":"10000"}],"bounds":[{"start":2853,"rate_min":"10000","rate_max":"30000"}]}'),
	('oasis1qry45a506cper3uslzemseek9qffg9qz5c46xrde', '{"rates":[{"start":1624,"rate":"4000"},{"start":1645,"rate":"2000"}],"bounds":[{"start":1624,"rate_min":"4000","rate_max":"10000"},{"start":1645,"rate_min":"2000","rate_max":"3000"}]}'),
	('oasis1qq7vyz4ewrdh00yujw0mgkf459et306xmvh2h3zg', '{"rates":[{"rate":"5000"},{"start":292,"rate":"10000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qrs8xpzhevja678g435afvkvs4h6xm6qzgdz6slf', '{"rates":[{"start":3147,"rate":"15000"}],"bounds":[{"start":3147,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qpp3q8rezyes0mk6ktavuyp8x5guvjmms59u7rda', '{"rates":[{"start":2510,"rate":"0"}],"bounds":[{"start":2510,"rate_min":"0","rate_max":"15000"}]}'),
	('oasis1qz26ty8q6gwt6zah7dtt8jpepvwnttkg8ssnxjl7', '{"rates":[{"rate":"5000"},{"start":4079,"rate":"20000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qr9wccuk6pqcr5ld8t2uf599e4ch5348hqeq53x4', '{"rates":[{"rate":"5000"},{"start":1476,"rate":"15000"},{"start":1587,"rate":"20000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":1476,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qql4alk30frfa6xua42eu7tynkqf9vd5ug95yqpn', '{"rates":[{"start":1228,"rate":"19500"},{"start":4643,"rate":"20000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qz04lk6qzm9ml3ff4qt3f7qjepdqh4l6dyr4lsmu', '{"rates":[{"start":5940,"rate":"10000"}],"bounds":[{"start":5940,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qq0lrzantpn5a4gx8ej7ampfw25cy8z78vh4uep7', '{"rates":[{"start":9234,"rate":"10000"},{"start":9601,"rate":"15000"}],"bounds":[{"start":9234,"rate_min":"0","rate_max":"25000"},{"start":9601,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qzm74el4utw4jssrl95ujq87g3ks3xfmjytvtaaa', '{"rates":[{"start":882,"rate":"20000"},{"start":898,"rate":"19000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qrs8zlh0mj37ug0jzlcykz808ylw93xwkvknm7yc', '{"rates":[{"start":3673,"rate":"19000"},{"start":11070,"rate":"20000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qp6fzgx9zhamsk6c77cwzjeme06xwswffvhk6js2', '{"rates":[{"start":485,"rate":"18000"},{"start":1212,"rate":"19000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qz72lvk2jchk0fjrz7u2swpazj3t5p0edsdv7sf8', '{"rates":[{"start":86,"rate":"15000"},{"start":1693,"rate":"19000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qrecdyjt5wu4ynkugv9ns8qjl5wl7gzprukywrly', '{"rates":[{"start":2221,"rate":"10000"}],"bounds":[{"start":2221,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzl8w5ka9y3p8a8gqlemqk98hzc33sn0tuezyc8l', '{"rates":[{"rate":"5000"},{"start":510,"rate":"10000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":510,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qpjuke27se2wnmvx6e8uc4l5h44yjp9h7g2clqfq', '{"rates":[{"rate":"5000"},{"start":2381,"rate":"20000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qrcmmvgf7qt8pmy543dmz0muwcm752kleund8rwr', '{"rates":[{"rate":"5000"},{"start":1260,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":1596,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qpw9clas6p53f44vq9vu7m5kzta2z23gcu8qezq5', '{"rates":[{"start":3210,"rate":"1000"},{"start":4450,"rate":"2000"}],"bounds":[{"start":3210,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qzzwkrjmyrzey46hyste5xszue3ggy90ag7k3687', '{"rates":[{"rate":"5000"},{"start":1259,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":1595,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qryreqam7w0slj7rhz70g9xvt9rct2024cepgqjj', '{"rates":[{"rate":"5000"},{"start":1259,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":1595,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qzws86jlwurt3e4vrm9tmywgpamhn7l8mglrxl6h', '{"rates":[{"rate":"5000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qrcaq6fgn9dnmg444asts3qkutur24hvpqp7luwa', '{"rates":[{"start":4724,"rate":"5000"}],"bounds":[{"start":4724,"rate_min":"5000","rate_max":"25000"}]}'),
	('oasis1qqewwznmvwfvee0dyq9g48acy0wcw890g549pukz', '{"rates":[{"start":26,"rate":"15000"},{"start":443,"rate":"19000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qqftj34pqk0kl5yz89zy86wvjqyytpwkfug04dgj', '{"rates":[{"start":2118,"rate":"7500"}],"bounds":[{"start":2118,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qra3rvq7y055waxmnx8rc0nad3frr8na2s9l8l3f', '{"rates":[{"start":52,"rate":"12500"},{"start":104,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qp0j5v5mkxk3eg4kxfdsk8tj6p22g4685qk76fw6', '{"rates":[{"start":5369,"rate":"5000"},{"start":9968,"rate":"20000"}],"bounds":[{"start":5369,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qqgdt8ajchz7gytdh7q2yfk0xmq93mhn7g0l2lje', '{"rates":[{"start":1599,"rate":"0"}],"bounds":[{"start":1599,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qpg763ex50jp0e34lsu789smlzqvcv7025v7m7yx', '{"rates":[{"rate":"5000"},{"start":1260,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":1596,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qpp6s22fa3l86enkye5wlu2jgp83zrgspg4fzlhr', '{"rates":[{"start":4064,"rate":"12000"}],"bounds":[{"start":4064,"rate_min":"0","rate_max":"100000"}]}'),
	('oasis1qz0tqva49ysnjk2p7xe83qfp86khxwms8sc2wf6e', '{"rates":[{"rate":"5000"},{"start":510,"rate":"10000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":510,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qzrehfnnntdaeshy5f6kfa8v3p35yu7mluaapmgc', '{"rates":[{"rate":"5000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qqqn56k79st0zy4060m0wvmec76vd3pl6czmg90a', '{"rates":[{"start":12981,"rate":"5000"}],"bounds":[{"start":12981,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qp3rhyfjagkj65cnn6lt8ej305gh3kamsvzspluq', '{"rates":[{"rate":"5000"},{"start":1758,"rate":"9000"}],"bounds":[{"rate_min":"0","rate_max":"20000"},{"start":1758,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzk6qlmgnq40cq2n3jfkw3307feqngt4gvksfml6', '{"rates":[{"start":10528,"rate":"20000"}],"bounds":[{"start":10528,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qqlq8jplttrmw982rrhp0svl8054xqevecks5yqr', '{"rates":[{"start":1730,"rate":"18000"},{"start":2144,"rate":"19500"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qrwxcth2c6njtyan2vdt7k9ra2xkmadmpv7v9387', '{"rates":[{"start":5635,"rate":"5000"},{"start":8720,"rate":"0"}],"bounds":[{"start":5635,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qrtvgssu2nxtnxdsgcgkmm2tj4pq9c0fy50ecazz', '{"rates":[{"start":12573,"rate":"100000"}],"bounds":[{"start":12573,"rate_min":"0","rate_max":"100000"}]}'),
	('oasis1qp53ud2pcmm73mlf4qywnrr245222mvlz5a2e5ty', '{"rates":[{"start":9366,"rate":"19000"},{"start":10227,"rate":"20000"}],"bounds":[{"start":9366,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qpqz2ut6a6prfcjm64xnpnjhsnqny6jqfyav829v', '{"rates":[{"start":7102,"rate":"1000"},{"start":11751,"rate":"0"}],"bounds":[{"start":7102,"rate_min":"0","rate_max":"20000"},{"start":11751,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzayckzyz35zsyan9zgv0gmhdq67ca8fhue6k4n8', '{"rates":[{"start":1754,"rate":"4000"}],"bounds":[{"start":1754,"rate_min":"0","rate_max":"4000"}]}'),
	('oasis1qzmwdlxy7cltmwt99u9pwqt3g0rdwgsqyvcqymmt', '{"rates":[{"start":418,"rate":"15000"},{"start":1220,"rate":"18000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qq7fn8e5zcl2q0x6t9a77ejgagt77anvhv3xlerw', '{"rates":[{"start":1599,"rate":"3000"}],"bounds":[{"start":1599,"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qz65awwegd9pr8msfxkg7hpwyjemm2qdlysyc8jq', '{"rates":[{"rate":"5000"},{"start":1646,"rate":"15000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qzrjsupecrjmvyfmg0aqz75sek4x7dn4ggchqx28', '{"rates":[{"start":1698,"rate":"10000"}],"bounds":[{"start":1698,"rate_min":"0","rate_max":"25000"}]}'),
	('oasis1qqdd4nmtcmarf4u9gdg24swxhec52du43cxzf302', '{"rates":[{"rate":"5000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}'),
	('oasis1qpgl52u29wy4hjla89f46ntkn2qsa6zpdvhv6s6n', '{"rates":[{"start":1641,"rate":"19000"}],"bounds":[{"start":1641,"rate_min":"19000","rate_max":"19500"}]}'),
	('oasis1qrtq873ddwnnjqyv66ezdc9ql2a07l37d5vae9k0', '{"rates":[{"start":28,"rate":"15000"},{"start":3148,"rate":"20000"}],"bounds":[{"rate_min":"0","rate_max":"20000"}]}');

COMMIT;