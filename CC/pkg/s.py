import random
import string

# List of common real domain names to base typos on
common_domains = [
    "google.com", "facebook.com", "youtube.com", "amazon.com", "yahoo.com",
    "wikipedia.org", "twitter.com", "instagram.com", "linkedin.com", "netflix.com",
    "ebay.com", "bing.com", "msn.com", "apple.com", "reddit.com"
]

# Typo generation functions
def swap_adjacent_chars(domain):
    for i in range(len(domain) - 1):
        if domain[i].isalpha() and domain[i+1].isalpha():
            return domain[:i] + domain[i+1] + domain[i] + domain[i+2:]
    return domain

def omit_char(domain):
    i = random.randint(0, len(domain) - 1)
    return domain[:i] + domain[i+1:]

def duplicate_char(domain):
    i = random.randint(0, len(domain) - 1)
    return domain[:i] + domain[i] + domain[i] + domain[i+1:]

def replace_char(domain):
    i = random.randint(0, len(domain) - 1)
    new_char = random.choice(string.ascii_lowercase)
    return domain[:i] + new_char + domain[i+1:]

# Combine all typo functions
typo_functions = [swap_adjacent_chars, omit_char, duplicate_char, replace_char]

# Generate 1000 fake domain names
fake_domains = set()
while len(fake_domains) < 1000:
    base = random.choice(common_domains)
    domain_part, tld = base.rsplit('.', 1)
    typo_func = random.choice(typo_functions)
    typo_domain = typo_func(domain_part) + '.' + tld
    fake_domains.add('"' + typo_domain + '"')
 
# Create a comma-separated string
fake_domains_str = ", ".join(sorted(fake_domains))
print(fake_domains_str)

