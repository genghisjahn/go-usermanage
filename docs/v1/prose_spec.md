## Go User Managerment Prose Specification

### Problem

I often think of software projects that I'd like to build/work on and many of these would require some form of user management.  By user management I mean the creation of user accounts (with something like email and password) and then using that email and password later to provide authentcation to a service.

However, I do not such a service at the momement.  I could use ones provided by various cloud providers, but the I'm beholden to those platforms for my users.  I'd rather have control.

I could use an already existing framework (there are many) but the ones I've looked at for Go (golang) seem complicated or use words/terms that I do not understand.  I'd end up implementing one of these frameworks poorly and fouling up the configuration.

So I'm going to write my own because:

1. I need it
1. It will be fun
1. Even if I can't get it to work, I will learn tremendous amounts so that should I cave in and using something off the shelf I'll have a much better understanding of how they work

### What will version 1 do?

1. Accept an email address and a password to create an account
    1. It will write those values to a database 
    1. It will add the bool columns `verified`  and `inactive`
        1. Verified will to make sure that the email address is valid (the user clicks a link)
        1. Inactive will be for accounts that are disabled for whatever reason
    1. We will validate around email address, at creation we will just validate that the user has entered something that is at least a single character followed by an `@` symbol and followed by at least one other character
        1.  A later version should include the ability to check for valid email addresses using a package that does this [this](https://www.interserver.net/tips/kb/check-email-address-really-exists-without-sending-email/).
    1. Return a verification GUID to be used to verify that the email address is correct
    1. It is up to the calling application to handle the sending of the email to the user containing the GUID
1. Accept a verification GUID verify that they exist and if they do, update the verified value in the users record
1. Accept an email address and password as login arguments, and the  check to see if the email address and password are valid, if so, return a Refresh Token

The Refresh token will be a JWT using the RSA signature.  It will be signed with a private key stored in the service.  

I had wanted this service to be multi-tennant but that will have to come later, for now I'll assume that each of service I build (assuming I get around to building them) will have it's own implementation.