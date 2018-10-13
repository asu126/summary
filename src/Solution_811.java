package leetcode;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Iterator;
import java.util.List;
import java.util.Map;

public class Solution_811 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_811 cc = new Solution_811();
		String[] s = { "900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org" };
		System.out.println(cc.subdomainVisits(s));

		// System.out.println("intel.mail.com".split("\\.")[0]);

	}

	public List<String> subdomainVisits(String[] cpdomains) {
		List<String> list = new ArrayList<String>();
		Map<String, Integer> map = new HashMap<String, Integer>();

		for (String s : cpdomains) {
			String[] tmp = s.split(" ");
			if (tmp.length != 2) {
				break;
			}

			String[] tmp1 = tmp[1].split("\\.");
			int value = Integer.parseInt(tmp[0]);

			if (tmp1.length >= 1) {
				StringBuilder sb = new StringBuilder().append(tmp1[tmp1.length - 1]);
				if (map.get(sb.toString()) != null) {
					map.put(sb.toString(), (map.get(sb.toString())) + value);
				} else {
					map.put(sb.toString(), value);
				}
				for (int i = tmp1.length - 2; i >= 0; i--) {
					// System.out.println(tmp1[i]);
					String t = tmp1[i];
					// sb.append(".").append(t);
					sb.insert(0, ".").insert(0, t);
					if (map.get(sb.toString()) != null) {
						map.put(sb.toString(), (map.get(sb.toString())) + value);
					} else {
						map.put(sb.toString(), value);
					}
				}
			}
		}

		// System.out.println(map);

		Iterator iterator = map.keySet().iterator();
		while (iterator.hasNext()) {
			String key = iterator.next().toString();
			list.add(map.get(key) + " " + key);
		}
		return list;
	}
}
