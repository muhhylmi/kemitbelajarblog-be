-- 002_seed_data.sql
-- Seeds the default authors and blog posts from the frontend

-- Insert authors (using deterministic UUIDs for consistency)
INSERT INTO authors (id, name, avatar, role) VALUES
(
    'a0000000-0000-0000-0000-000000000001',
    'Elias Thorne',
    'https://lh3.googleusercontent.com/aida-public/AB6AXuD4VSMhOs2fuadCiq341O5vuYntg34QmmQLIaJ17eUOqlqVDC_aHFnJIcI_5VQ18VteftA-Ie_MZ8ocu4uN02FCgKlJsZMPu3bbhQ7ltdY_x3JKb8snsfLzPtcrHOGV5MbJIBMsjxQsiIqsKt_yykuijuB0TSXQ5Tg6GT-Yxrxc-2MUO5vTJwhTVl_GOZURV95HDr_G1BHfcJgmNFCO8pBNxYWTzz2Rk2YQX1yyGw3IhaFPq40kwDX3uP-SorptmSX1CuWjcaDqj7w',
    'Author'
),
(
    'a0000000-0000-0000-0000-000000000002',
    'Julian Thorne',
    'https://lh3.googleusercontent.com/aida-public/AB6AXuBNeAwWXipgYw3ebDf59V9RNm-uMN4KS-xegLQ6MO03vINNk-Lohg5ZvNDWZfmVjrWd6CjZjxT_pkfy85BTqjgbrfgMqYAC-Qc1LeofR06coAjGqJYJanoALZ8rs2l3-USqLAFr6huXS0i0ny5JwRVbd2N11SneO_OSqMH0OP9mQvMCr566w-Fqc5yavvfxomVWLnBWIfAHaeRJ_lHbqW1eAd5_--mQ8Hkqkni_e4e3iOVX5YUYqbWozHssYkXxbpPb4_Lohrzl9NM',
    'Senior Editor'
)
ON CONFLICT (id) DO NOTHING;

-- Insert default blog posts
INSERT INTO posts (id, title, summary, content, category, author_id, read_time, published_at, image, image_alt, likes, is_featured, status, created_at) VALUES
(
    'philosophy-slow-reading',
    'The Philosophy of Slow Reading in an Instant World',
    'How returning to tactile and deliberate engagement with text can reshape our cognitive landscape and emotional well-being.',
    E'In our current digital epoch, speed is often confused with progress. We move from task to task, notification to notification, with a kinetic energy that leaves little room for the quiet resonance of the present. Rooted living is not a rejection of technology, but a recalibration of our relationship with it. It is the intentional choice to anchor ourselves in the physical, the tactile, and the enduring.\n\n## The Slow Observation\n\nNature does not hurry, yet everything is accomplished. When we observe a forest, we see a complex network of interdependence that thrives on patience. The trees do not compete for the sky in a way that exhausts the soil; they grow in harmony with their environment. This is the first principle of rooted living: environmental awareness.\n\nOur domestic spaces often mirror our internal states. A cluttered, sterile environment breeds a cluttered, anxious mind. By introducing natural textures—wood, linen, stone—and earthy tones into our homes, we create a sensory bridge back to the earth. These materials carry a weight and a history that plastic simply cannot replicate.\n\n## Cultivating Presence\n\nThe practice of deep reading is perhaps one of the most effective ways to cultivate this rootedness. Unlike the fragmented skimming we perform on social feeds, deep reading requires a sustained attention that rewires our neurocircuitry for depth. It is a form of secular meditation, where the author''s voice becomes a companion in the silence.\n\nAs we move forward into an increasingly automated future, the human need for the "organic" will only intensify. Rooted living is our way of ensuring that while we may reach for the stars through our technology, our feet remain firmly planted in the soil that sustains us.',
    'Featured Insight',
    'a0000000-0000-0000-0000-000000000001',
    '12 min read',
    'Today',
    'https://lh3.googleusercontent.com/aida-public/AB6AXuCRQk2D-RJ_T3uFzWmfnWzxuJjlnbvZQIKJ1qy3DMOC0_mtKMbJoK1meDX07Ub_pFceTD-mM2aVXAk6evo6_AiHzf7mbqElbOQvqkXuYMNemjDiyYIH48IVGnbth81H0UykddIVbCgyA5-k19bQE2kGDIO-yNGRvPfRVq3yTie-txrNHVlIrssgEN2slx4qToA3KdCrDJ29Z-BJe-sOEWk8zYTWkssho1j8IzrtPdeFwhDsa4KhX6Yf3qOSnN3JAuWo8gX8ltHbdAw',
    'A lush, sun-drenched forest clearing where golden light filters through ancient mossy trees.',
    1200,
    TRUE,
    'published',
    NOW() - INTERVAL '1 hour'
),
(
    'quietening-digital-noise',
    'Quietening the Digital Noise',
    'Finding pockets of silence in a world that never stops talking.',
    E'Finding pockets of silence in a world that never stops talking is more than a convenience—it is a necessity for mental clarity.\n\n## The Overwhelm of Connectivity\n\nEvery day, we are bombarded with notifications, emails, alerts, and feeds. This constant flow of information puts our nervous system on high alert. We are always reacting, rarely reflecting.\n\n## Crafting Pockets of Silence\n\nTo find quiet, we must intentionally design it into our day. Here are three simple techniques:\n\n1. **Morning Silence**: Do not touch your phone for the first 30 minutes of the day. Sip your tea, look out the window, and let your thoughts wake up naturally.\n2. **Device-Free Zones**: Designate your dining table or bedroom as a screen-free zone.\n3. **Tactile Alternations**: Trade one hour of digital media for a physical book, a sketchpad, or gardening.\n\nBy silencing the noise, we allow our inner voice to speak.',
    'Mindfulness',
    'a0000000-0000-0000-0000-000000000002',
    '8 min read',
    '2 days ago',
    'https://lh3.googleusercontent.com/aida-public/AB6AXuBdjeGIbDgGh1BZIIb0bocdNMrzDLkeaYFtBaqNnff3Ib1TlZmOFgyNu1MsRLEH_3pXD_rkFm3gnBHd1dpcJKttGm5wYZhErVi109nnpZkWKv4vHWFkce4GrgIxsnCR966qpRyhJLJoV7PoXJV45GazHAc4E5rcx8cu8cPcsDznf3mf6p-R1VIfOVmL0cv8ADHdgcl2NRI_mgYjGd85ZcFF3tBFqJ2y2O4dI9Bedw7NvzXMJJSma3ZmouCReUD0q1p-B_hXSjiNtJ0',
    'A peaceful sunset over a calm lake with silhouettes of reeds in the foreground.',
    840,
    FALSE,
    'published',
    NOW() - INTERVAL '2 days'
),
(
    'rooted-living',
    'The Art of Rooted Living: Finding Stillness in a Kinetic World',
    'How returning to tactile and deliberate engagement with text can reshape our cognitive landscape and emotional well-being.',
    E'"To be rooted is perhaps the most important and least recognized need of the human soul. It is one of the hardest to define." — Simone Weil\n\nIn our current digital epoch, speed is often confused with progress. We move from task to task, notification to notification, with a kinetic energy that leaves little room for the quiet resonance of the present. Rooted living is not a rejection of technology, but a recalibration of our relationship with it. It is the intentional choice to anchor ourselves in the physical, the tactile, and the enduring.\n\n## The Slow Observation\n\nNature does not hurry, yet everything is accomplished. When we observe a forest, we see a complex network of interdependence that thrives on patience. The trees do not compete for the sky in a way that exhausts the soil; they grow in harmony with their environment. This is the first principle of rooted living: environmental awareness.\n\nOur domestic spaces often mirror our internal states. A cluttered, sterile environment breeds a cluttered, anxious mind. By introducing natural textures—wood, linen, stone—and earthy tones into our homes, we create a sensory bridge back to the earth. These materials carry a weight and a history that plastic simply cannot replicate.\n\n## Cultivating Presence\n\nThe practice of deep reading is perhaps one of the most effective ways to cultivate this rootedness. Unlike the fragmented skimming we perform on social feeds, deep reading requires a sustained attention that rewires our neurocircuitry for depth. It is a form of secular meditation, where the author''s voice becomes a companion in the silence.\n\nAs we move forward into an increasingly automated future, the human need for the "organic" will only intensify. Rooted living is our way of ensuring that while we may reach for the stars through our technology, our feet remain firmly planted in the soil that sustains us.',
    'Sustainability',
    'a0000000-0000-0000-0000-000000000002',
    '12 min read',
    'October 12, 2023',
    'https://lh3.googleusercontent.com/aida-public/AB6AXuDPAl77cDYxebbWSZ4sIpQc6IyVjgLR_GtmILx1Gj5REZwh2Xd--CW-WSiAAWM4SMZ4tSHjgScP8_W8Wv7qMkm-1IFV9nyW18_PDiSBpXADhu_r-CNNXBKNW5vMEw3t8fQsw8F4CGTh9gLWfm-ZmHzA4qsG83dVSmYFCA0MoFrGzUkO7eWNsekYtD7jMDkYiWciUO1pJa5Uaht4exIn1GqW6U30eXyDLebeFkhM7l_n4N2BjnBEW_8cgAGn7pd4qhrlTOZZMkVP30g',
    'A lush, serene forest scene featuring a single, majestic old-growth tree with deep roots spreading across a moss-covered ground.',
    1200,
    FALSE,
    'published',
    NOW() - INTERVAL '7 days'
),
(
    'intelligence-old-growth',
    'The Intelligence of Old Growth',
    'What ancient root systems can teach us about community resilience and growth.',
    E'Deep beneath the forest floor lies a complex network of roots and mycelium linking individual trees in an ancient web of mutual aid.\n\n## The Mycorrhizal Web\n\nTrees are not isolated individuals competing for light; they are part of a cooperative society. Through mycorrhizal fungi, older "mother trees" share nutrients and water with struggling saplings, warn neighbors of pests, and coordinate responses to environmental stresses.\n\n## Lessons for Community Resilience\n\n1. **Cooperation Over Competition**: Growth is more sustainable when we support the collective health of our community rather than chasing individual gain.\n2. **Slow Accumulation**: The resilience of old growth comes from decades of slow, steady adaptation. Speed often builds fragility.',
    'Nature',
    'a0000000-0000-0000-0000-000000000002',
    '15 min read',
    '5 days ago',
    'https://lh3.googleusercontent.com/aida-public/AB6AXuCqKFzgITbsk8bk9ttQmOO14JmqJVomZ2h3_bicQ9mM5mGbaPKJRCbQV6ztedIQtdES3L9sJs3uB6QturRPNVCMApdFIHpisgndPrBW3aYEtaIdgNFwceb12TndCDPzyZ_H7zWkYOTcidqWyYxCeD8OShxcL6odTes4Y8VitZOuLuLRnbZjsdOJ3x6WvHcWB0OAEJsSygzg2bfCwh9Vfs7i6P3NDgolDRLiAfmg6Gv34L7ZoWuRfjFjnuNzlFDePq2XmA8AqjMIvzo',
    'Magnificent mountain peaks reaching into a soft, hazy sky during the golden hour.',
    512,
    FALSE,
    'published',
    NOW() - INTERVAL '5 days'
),
(
    'manual-craftsmanship',
    'Manual Craftsmanship',
    'Understanding the tactile connection between the maker, the material, and the finished work in a digital age.',
    E'There is a deep cognitive and emotional satisfaction in making things by hand. From throwing clay on a wheel to woodworking, manual craftsmanship links us directly to the physical world.\n\n## The Feedback Loop of Matter\n\nWhen working with clay, wood, or metal, the material responds instantly to your touch. It has limits, grain, and resistance. You cannot override these boundaries with a shortcut; you must listen to the material and adapt. This reciprocal feedback loop creates a flow state that is rare in digital interfaces.\n\n## Restoring the Hand-Brain Connection\n\nStudies show that using our hands to shape three-dimensional objects builds strong neuro-pathways linked to spatial awareness, problem-solving, and emotional grounding. Craftsmanship is not just a hobby; it is a vital counterweight to our digital lives.',
    'Manual Craftsmanship',
    'a0000000-0000-0000-0000-000000000001',
    '5 min read',
    '1 week ago',
    'https://lh3.googleusercontent.com/aida-public/AB6AXuB6CDKgF3QQpV8wK-iKq27BjzWiSrWtvH2DUSbHEiDkPUING0vMm_eqNutxbG3Ewq7TjjJSqQBpbdz0F4P1vIX8uugfFfEHLZOYgkp4eXdFjK3sFrXfKA_-QJ7vkuzVvIYZKT2A2rIjvhSUjRO4KLbRJ3Gb87bh10OJQf3ALwT3AwqXSuUSZ58tMqvdb68AIQSmQv6e87DMDChZM9e4ca0UukTBj3peV46r_OkRG8Wt__WTIkCqIp6gxZw0db5qutcZ_XMc-Y-Yt1I',
    'A macro photograph of ceramic pottery with an organic, uneven glaze in earthy terracotta and sage tones.',
    320,
    FALSE,
    'published',
    NOW() - INTERVAL '1 week'
),
(
    'drafting-messy-middle',
    'Drafting: A Guide to the Messy Middle',
    'Why the first draft is for the writer, and the second is for the reader. Tips on overcoming the creative slump.',
    E'The blank page is intimidating, but the middle of a project is where most creative endeavors stall. Here is how to navigate the drafting process.\n\n## The First Draft is for You\n\nWhen writing the first draft, silence your inner critic. Do not worry about spelling, grammar, or flow. The sole purpose of the first draft is to transfer your thoughts from your mind to the page. You cannot edit what does not exist.\n\n## The Second Draft is for the Reader\n\nOnce you have the raw material, switch roles. Put on your editor hat. Ask yourself:\n\n- Is this clear?\n- Does the structure flow logically?\n- Can I cut filler words?\n\nEditing is the art of subtraction. By peeling away the extra layers, the true essence of your ideas will shine through.',
    'Writing',
    'a0000000-0000-0000-0000-000000000002',
    '20 min read',
    '2 weeks ago',
    '',
    '',
    195,
    FALSE,
    'draft',
    NOW() - INTERVAL '2 weeks'
)
ON CONFLICT (id) DO NOTHING;
