U
    2�^�  �                   @   s�  d dl mZmZmZmZmZ d dlmZ d dlZd dl	Z	ddl
T d dl mZ d dlmZ G dd	� d	e�ZG d
d� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG d d!� d!e�ZG d"d#� d#e�ZG d$d%� d%e�ZG d&d'� d'e�ZG d(d)� d)e�ZG d*d+� d+e�Z G d,d-� d-e�Z!G d.d/� d/e�Z"G d0d1� d1e�Z#G d2d3� d3e�Z$G d4d5� d5e�Z%G d6d7� d7e�Z&G d8d9� d9e�Z'dS ):�    )�TType�TMessageType�TFrozenDict�
TException�TApplicationException)�TProtocolExceptionN�   )�*)�
TProcessor)�
TTransportc                	   @   s|   e Zd ZdZdddejdddfdejdddfdejd	ddffZdd
d�Zdd� Zdd� Z	dd� Z
dd� Zdd� Zdd� ZdS )�normalizePhoneNumber_argszS
    Attributes:
     - countryCode
     - phoneNumber
     - countryCodeHint

    N�   �countryCode�UTF8�   �phoneNumber�   �countryCodeHintc                 C   s   || _ || _|| _d S �N)r   r   r   ��selfr   r   r   � r   �//storage/sdcard0/bot1/flex/curve/AuthService.py�__init__   s    z"normalizePhoneNumber_args.__init__c                 C   sX  |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr`�qL|dkr�|t	jkr�tjd dkr�|�� �d�n|�� | _n
|�|� n�|dkr�|t	jkr�tjd dkr�|�� �d�n|�� | _n
|�|� nV|dk�r8|t	jk�r,tjd dk�r |�� �d�n|�� | _n
|�|� n
|�|� |��  qD|��  d S )Nr   r   �utf-8r   r   )�_fast_decode�
isinstance�transr   �CReadableTransport�thrift_spec�	__class__�readStructBegin�readFieldBeginr   �STOP�STRING�sys�version_info�
readString�decoder   �skipr   r   �readFieldEnd�readStructEnd�r   �iprot�fname�ftype�fidr   r   r   �read!   s,    "

(
(
*

znormalizePhoneNumber_args.readc                 C   s.  |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  | jd k	r�|�dtj	d� |�
tjd dkr�| j�d�n| j� |��  | jd k	�r|�dtj	d	� |�
tjd dk�r
| j�d�n| j� |��  |��  |��  d S )
Nr   r   r   r   r   r   r   r   r   )�_fast_encoder   r   �writer    �writeStructBeginr   �writeFieldBeginr   r$   �writeStringr%   r&   �encode�writeFieldEndr   r   �writeFieldStop�writeStructEnd�r   �oprotr   r   r   r3   >   s$    

&
&(znormalizePhoneNumber_args.writec                 C   s   d S r   r   �r   r   r   r   �validateR   s    z"normalizePhoneNumber_args.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS �z%s=%rr   ��.0�key�valuer   r   r   �
<listcomp>V   s   �z6normalizePhoneNumber_args.__repr__.<locals>.<listcomp>�%s(%s)�, ��__dict__�itemsr    �__name__�join�r   �Lr   r   r   �__repr__U   s    �z"normalizePhoneNumber_args.__repr__c                 C   s   t || j�o| j|jkS r   �r   r    rH   �r   �otherr   r   r   �__eq__Z   s    z normalizePhoneNumber_args.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   �__ne__]   s    z normalizePhoneNumber_args.__ne__)NNN�rJ   �
__module__�__qualname__�__doc__r   r$   r   r   r1   r3   r>   rN   rR   rS   r   r   r   r   r      s   �
r   c                   @   sn   e Zd ZdZdejdddfdejdedgdffZddd	�Z	d
d� Z
dd� Zdd� Zdd� Zdd� Zdd� ZdS )�normalizePhoneNumber_result�.
    Attributes:
     - success
     - e

    r   �successr   Nr   �ec                 C   s   || _ || _d S r   �rZ   r[   �r   rZ   r[   r   r   r   r   l   s    z$normalizePhoneNumber_result.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�tjd dkr�|�� �d�n|�� | _q�|�|� n>|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S �Nr   r   r   r   �r   r   r   r   r   r   r    r!   r"   r   r#   r$   r%   r&   r'   r(   rZ   r)   �STRUCT�TalkExceptionr[   r1   r*   r+   r,   r   r   r   r1   p   s&    "

(


z normalizePhoneNumber_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  | jd k	r�|�dtjd� | j�|� |��  |��  |��  d S )NrX   rZ   r   r   r   r[   r   �r2   r   r   r3   r    r4   rZ   r5   r   r$   r6   r%   r&   r7   r8   r[   r`   r9   r:   r;   r   r   r   r3   �   s    

&
z!normalizePhoneNumber_result.writec                 C   s   d S r   r   r=   r   r   r   r>   �   s    z$normalizePhoneNumber_result.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   �   s   �z8normalizePhoneNumber_result.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   �   s    �z$normalizePhoneNumber_result.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   �   s    z"normalizePhoneNumber_result.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   �   s    z"normalizePhoneNumber_result.__ne__)NN�rJ   rU   rV   rW   r   r$   r`   ra   r   r   r1   r3   r>   rN   rR   rS   r   r   r   r   rX   `   s   �
rX   c                
   @   s�   e Zd ZdZddejdddfdejdedgdfdejd	d
dfdejdd
dfdejdddffZ	ddd�Z
dd� Zdd� Zdd� Zdd� Zdd� Zdd� ZdS )�respondE2EELoginRequest_argszu
    Attributes:
     - verifier
     - publicKey
     - encryptedKeyChain
     - hashKeyChain
     - errorCode

    Nr   �verifierr   r   �	publicKeyr   �encryptedKeyChain�BINARYr   �hashKeyChain�   �	errorCodec                 C   s"   || _ || _|| _|| _|| _d S r   )re   rf   rg   ri   rk   �r   re   rf   rg   ri   rk   r   r   r   r   �   s
    z%respondE2EELoginRequest_args.__init__c                 C   s�  |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr`�qv|dkr�|t	jkr�tjd dkr�|�� �d�n|�� | _n
|�|� n�|dkr�|t	jkr�t� | _| j�|� n
|�|� n�|dk�r|t	jkr�|�� | _n
|�|� nf|dk�r4|t	jk�r(|�� | _n
|�|� n8|dk�rb|t	jk�rV|�� | _n
|�|� n
|�|� |��  qD|��  d S )Nr   r   r   r   r   r   rj   )r   r   r   r   r   r   r    r!   r"   r   r#   r$   r%   r&   r'   r(   re   r)   r`   �E2EEPublicKeyrf   r1   �
readBinaryrg   ri   �I32�readI32rk   r*   r+   r,   r   r   r   r1   �   s>    "

(






z!respondE2EELoginRequest_args.readc                 C   sV  |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  | jd k	r�|�dtjd� | j�|� |��  | jd k	r�|�dtj	d	� |�| j� |��  | jd k	�r|�d
tj	d� |�| j� |��  | jd k	�rB|�dtjd� |�| j� |��  |��  |��  d S )Nrd   re   r   r   r   r   rf   rg   r   ri   r   rk   rj   )r2   r   r   r3   r    r4   re   r5   r   r$   r6   r%   r&   r7   r8   rf   r`   rg   �writeBinaryri   rk   ro   �writeI32r9   r:   r;   r   r   r   r3   �   s4    

&

z"respondE2EELoginRequest_args.writec                 C   s   d S r   r   r=   r   r   r   r>     s    z%respondE2EELoginRequest_args.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD     s   �z9respondE2EELoginRequest_args.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN     s    �z%respondE2EELoginRequest_args.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR     s    z#respondE2EELoginRequest_args.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS     s    z#respondE2EELoginRequest_args.__ne__)NNNNN)rJ   rU   rV   rW   r   r$   r`   rm   ro   r   r   r1   r3   r>   rN   rR   rS   r   r   r   r   rd   �   s   
�
(rd   c                   @   sb   e Zd ZdZddejdedgdffZddd�Zdd� Z	d	d
� Z
dd� Zdd� Zdd� Zdd� ZdS )�respondE2EELoginRequest_result�
    Attributes:
     - e

    Nr   r[   c                 C   s
   || _ d S r   �r[   �r   r[   r   r   r   r     s    z'respondE2EELoginRequest_result.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S �Nr   �r   r   r   r   r   r   r    r!   r"   r   r#   r`   ra   r[   r1   r)   r*   r+   r,   r   r   r   r1      s    "



z#respondE2EELoginRequest_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  |��  |��  d S )Nrs   r[   r   �r2   r   r   r3   r    r4   r[   r5   r   r`   r8   r9   r:   r;   r   r   r   r3   4  s    

z$respondE2EELoginRequest_result.writec                 C   s   d S r   r   r=   r   r   r   r>   @  s    z'respondE2EELoginRequest_result.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   D  s   �z;respondE2EELoginRequest_result.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   C  s    �z'respondE2EELoginRequest_result.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   H  s    z%respondE2EELoginRequest_result.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   K  s    z%respondE2EELoginRequest_result.__ne__)N�rJ   rU   rV   rW   r   r`   ra   r   r   r1   r3   r>   rN   rR   rS   r   r   r   r   rs     s   �
rs   c                   @   sn   e Zd ZdZdddejdddfdejdddffZddd	�Zd
d� Z	dd� Z
dd� Zdd� Zdd� Zdd� ZdS )�getAuthRSAKey_argszC
    Attributes:
     - authSessionId
     - identityProvider

    Nr   �authSessionIdr   r   �identityProviderc                 C   s   || _ || _d S r   )r|   r}   �r   r|   r}   r   r   r   r   \  s    zgetAuthRSAKey_args.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�tjd dkr�|�� �d�n|�� | _q�|�|� n4|dkr�|t	jkr�|�� | _q�|�|� n
|�|� |��  qD|��  d S )Nr   r   r   r   )r   r   r   r   r   r   r    r!   r"   r   r#   r$   r%   r&   r'   r(   r|   r)   ro   rp   r}   r*   r+   r,   r   r   r   r1   `  s$    "

(


zgetAuthRSAKey_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  | jd k	r�|�dtjd� |�| j� |��  |��  |��  d S )Nr{   r|   r   r   r   r}   r   )r2   r   r   r3   r    r4   r|   r5   r   r$   r6   r%   r&   r7   r8   r}   ro   rr   r9   r:   r;   r   r   r   r3   x  s    

&
zgetAuthRSAKey_args.writec                 C   s   d S r   r   r=   r   r   r   r>   �  s    zgetAuthRSAKey_args.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   �  s   �z/getAuthRSAKey_args.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   �  s    �zgetAuthRSAKey_args.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   �  s    zgetAuthRSAKey_args.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   �  s    zgetAuthRSAKey_args.__ne__)NN)rJ   rU   rV   rW   r   r$   ro   r   r   r1   r3   r>   rN   rR   rS   r   r   r   r   r{   N  s   �
r{   c                   @   sb   e Zd ZdZddejdedgdffZddd�Zdd� Z	d	d
� Z
dd� Zdd� Zdd� Zdd� ZdS )�getAuthRSAKey_resultrt   Nr   r[   c                 C   s
   || _ d S r   ru   rv   r   r   r   r   �  s    zgetAuthRSAKey_result.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S rw   rx   r,   r   r   r   r1   �  s    "



zgetAuthRSAKey_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr   r[   r   ry   r;   r   r   r   r3   �  s    

zgetAuthRSAKey_result.writec                 C   s   d S r   r   r=   r   r   r   r>   �  s    zgetAuthRSAKey_result.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   �  s   �z1getAuthRSAKey_result.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   �  s    �zgetAuthRSAKey_result.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   �  s    zgetAuthRSAKey_result.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   �  s    zgetAuthRSAKey_result.__ne__)Nrz   r   r   r   r   r   �  s   �
r   c                   @   sl   e Zd ZdZddejdddfdejdddffZdd	d
�Zdd� Zdd� Z	dd� Z
dd� Zdd� Zdd� ZdS )�confirmE2EELogin_argsz:
    Attributes:
     - verifier
     - deviceSecret

    Nr   re   r   r   �deviceSecretrh   c                 C   s   || _ || _d S r   )re   r�   �r   re   r�   r   r   r   r   �  s    zconfirmE2EELogin_args.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�tjd dkr�|�� �d�n|�� | _q�|�|� n4|dkr�|t	jkr�|�� | _q�|�|� n
|�|� |��  qD|��  d S )Nr   r   r   r   )r   r   r   r   r   r   r    r!   r"   r   r#   r$   r%   r&   r'   r(   re   r)   rn   r�   r*   r+   r,   r   r   r   r1   �  s$    "

(


zconfirmE2EELogin_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  | jd k	r�|�dtj	d� |�| j� |��  |��  |��  d S )Nr�   re   r   r   r   r   r�   )r2   r   r   r3   r    r4   re   r5   r   r$   r6   r%   r&   r7   r8   r�   rq   r9   r:   r;   r   r   r   r3   �  s    

&
zconfirmE2EELogin_args.writec                 C   s   d S r   r   r=   r   r   r   r>     s    zconfirmE2EELogin_args.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD     s   �z2confirmE2EELogin_args.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN     s    �zconfirmE2EELogin_args.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR     s    zconfirmE2EELogin_args.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS     s    zconfirmE2EELogin_args.__ne__)NNrT   r   r   r   r   r�   �  s   �
r�   c                   @   sn   e Zd ZdZdejdddfdejdedgdffZddd	�Z	d
d� Z
dd� Zdd� Zdd� Zdd� Zdd� ZdS )�confirmE2EELogin_resultrY   r   rZ   r   Nr   r[   c                 C   s   || _ || _d S r   r\   r]   r   r   r   r   %  s    z confirmE2EELogin_result.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�tjd dkr�|�� �d�n|�� | _q�|�|� n>|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S r^   r_   r,   r   r   r   r1   )  s&    "

(


zconfirmE2EELogin_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  | jd k	r�|�dtjd� | j�|� |��  |��  |��  d S )Nr�   rZ   r   r   r   r[   r   rb   r;   r   r   r   r3   B  s    

&
zconfirmE2EELogin_result.writec                 C   s   d S r   r   r=   r   r   r   r>   R  s    z confirmE2EELogin_result.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   V  s   �z4confirmE2EELogin_result.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   U  s    �z confirmE2EELogin_result.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   Z  s    zconfirmE2EELogin_result.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   ]  s    zconfirmE2EELogin_result.__ne__)NNrc   r   r   r   r   r�     s   �
r�   c                   @   s`   e Zd ZdZdddejdddffZddd�Zdd� Zd	d
� Z	dd� Z
dd� Zdd� Zdd� ZdS )�*issueTokenForAccountMigrationSettings_argsz%
    Attributes:
     - enforce

    Nr   �enforcec                 C   s
   || _ d S r   )r�   �r   r�   r   r   r   r   l  s    z3issueTokenForAccountMigrationSettings_args.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr||�� | _q�|�|� n
|�|� |��  qD|��  d S �Nr   )r   r   r   r   r   r   r    r!   r"   r   r#   �BOOL�readBoolr�   r)   r*   r+   r,   r   r   r   r1   o  s    "



z/issueTokenForAccountMigrationSettings_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� |�
| j� |��  |��  |��  d S )Nr�   r�   r   )r2   r   r   r3   r    r4   r�   r5   r   r�   �	writeBoolr8   r9   r:   r;   r   r   r   r3   �  s    

z0issueTokenForAccountMigrationSettings_args.writec                 C   s   d S r   r   r=   r   r   r   r>   �  s    z3issueTokenForAccountMigrationSettings_args.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   �  s   �zGissueTokenForAccountMigrationSettings_args.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   �  s    �z3issueTokenForAccountMigrationSettings_args.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   �  s    z1issueTokenForAccountMigrationSettings_args.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   �  s    z1issueTokenForAccountMigrationSettings_args.__ne__)N)rJ   rU   rV   rW   r   r�   r   r   r1   r3   r>   rN   rR   rS   r   r   r   r   r�   `  s   �
r�   c                   @   sr   e Zd ZdZdejdedgdfdejdedgdffZddd�Z	d	d
� Z
dd� Zdd� Zdd� Zdd� Zdd� ZdS )�,issueTokenForAccountMigrationSettings_resultrY   r   rZ   Nr   r[   c                 C   s   || _ || _d S r   r\   r]   r   r   r   r   �  s    z5issueTokenForAccountMigrationSettings_result.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n>|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S �Nr   r   �r   r   r   r   r   r   r    r!   r"   r   r#   r`   �SecurityCenterResultrZ   r1   r)   ra   r[   r*   r+   r,   r   r   r   r1   �  s(    "




z1issueTokenForAccountMigrationSettings_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr�   rZ   r   r[   r   �r2   r   r   r3   r    r4   rZ   r5   r   r`   r8   r[   r9   r:   r;   r   r   r   r3   �  s    


z2issueTokenForAccountMigrationSettings_result.writec                 C   s   d S r   r   r=   r   r   r   r>   �  s    z5issueTokenForAccountMigrationSettings_result.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   �  s   �zIissueTokenForAccountMigrationSettings_result.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   �  s    �z5issueTokenForAccountMigrationSettings_result.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   �  s    z3issueTokenForAccountMigrationSettings_result.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   �  s    z3issueTokenForAccountMigrationSettings_result.__ne__)NN�rJ   rU   rV   rW   r   r`   r�   ra   r   r   r1   r3   r>   rN   rR   rS   r   r   r   r   r�   �  s   �
r�   c                   @   s`   e Zd ZdZdddejdddffZddd�Zdd	� Zd
d� Z	dd� Z
dd� Zdd� Zdd� ZdS )�"issueTokenForAccountMigration_argsz0
    Attributes:
     - migrationSessionId

    Nr   �migrationSessionIdr   c                 C   s
   || _ d S r   )r�   �r   r�   r   r   r   r   �  s    z+issueTokenForAccountMigration_args.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�tjd dkr�|�� �d�n|�� | _q�|�|� n
|�|� |��  qD|��  d S )Nr   r   r   )r   r   r   r   r   r   r    r!   r"   r   r#   r$   r%   r&   r'   r(   r�   r)   r*   r+   r,   r   r   r   r1   �  s    "

(

z'issueTokenForAccountMigration_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  |��  |��  d S )Nr�   r�   r   r   r   )r2   r   r   r3   r    r4   r�   r5   r   r$   r6   r%   r&   r7   r8   r9   r:   r;   r   r   r   r3     s    

&z(issueTokenForAccountMigration_args.writec                 C   s   d S r   r   r=   r   r   r   r>     s    z+issueTokenForAccountMigration_args.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD     s   �z?issueTokenForAccountMigration_args.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN     s    �z+issueTokenForAccountMigration_args.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR     s    z)issueTokenForAccountMigration_args.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS     s    z)issueTokenForAccountMigration_args.__ne__)NrT   r   r   r   r   r�   �  s   �
r�   c                   @   sr   e Zd ZdZdejdedgdfdejdedgdffZddd�Z	d	d
� Z
dd� Zdd� Zdd� Zdd� Zdd� ZdS )�$issueTokenForAccountMigration_resultrY   r   rZ   Nr   r[   c                 C   s   || _ || _d S r   r\   r]   r   r   r   r   ,  s    z-issueTokenForAccountMigration_result.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n>|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S r�   r�   r,   r   r   r   r1   0  s(    "




z)issueTokenForAccountMigration_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr�   rZ   r   r[   r   r�   r;   r   r   r   r3   J  s    


z*issueTokenForAccountMigration_result.writec                 C   s   d S r   r   r=   r   r   r   r>   Z  s    z-issueTokenForAccountMigration_result.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   ^  s   �zAissueTokenForAccountMigration_result.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   ]  s    �z-issueTokenForAccountMigration_result.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   b  s    z+issueTokenForAccountMigration_result.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   e  s    z+issueTokenForAccountMigration_result.__ne__)NNr�   r   r   r   r   r�      s   �
r�   c                   @   s�   e Zd ZdZdddejdddfdejdddfdejd	ddfd
ejdedgdfdejdddfdejdddffZ	ddd�Z
dd� Zdd� Zdd� Zdd� Zdd� Zdd� ZdS ) �verifyQrcodeWithE2EE_argsz�
    Attributes:
     - verifier
     - pinCode
     - errorCode
     - publicKey
     - encryptedKeyChain
     - hashKeyChain

    Nr   re   r   r   �pinCoder   rk   rj   rf   �   rg   rh   �   ri   c                 C   s(   || _ || _|| _|| _|| _|| _d S r   )re   r�   rk   rf   rg   ri   �r   re   r�   rk   rf   rg   ri   r   r   r   r   ~  s    z"verifyQrcodeWithE2EE_args.__init__c                 C   s�  |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr`�q�|dkr�|t	jkr�tjd dkr�|�� �d�n|�� | _n
|�|� �n|dkr�|t	jkr�tjd dkr�|�� �d�n|�� | _n
|�|� n�|dk�r|t	jk�r|�� | _n
|�|� n�|dk�rT|t	jk�rHt� | _| j�|� n
|�|� nf|dk�r�|t	jk�rv|�� | _n
|�|� n8|dk�r�|t	jk�r�|�� | _n
|�|� n
|�|� |��  qD|��  d S )	Nr   r   r   r   r   rj   r�   r�   )r   r   r   r   r   r   r    r!   r"   r   r#   r$   r%   r&   r'   r(   re   r)   r�   ro   rp   rk   r`   rm   rf   r1   rn   rg   ri   r*   r+   r,   r   r   r   r1   �  sF    "

(
(





zverifyQrcodeWithE2EE_args.readc                 C   s�  |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  | jd k	r�|�dtj	d� |�
tjd dkr�| j�d�n| j� |��  | jd k	r�|�dtjd	� |�| j� |��  | jd k	�r,|�d
tjd� | j�|� |��  | jd k	�r\|�dtj	d� |�| j� |��  | jd k	�r�|�dtj	d� |�| j� |��  |��  |��  d S )Nr�   re   r   r   r   r�   r   rk   r   rf   rj   rg   r�   ri   r�   )r2   r   r   r3   r    r4   re   r5   r   r$   r6   r%   r&   r7   r8   r�   rk   ro   rr   rf   r`   rg   rq   ri   r9   r:   r;   r   r   r   r3   �  s<    

&
&
zverifyQrcodeWithE2EE_args.writec                 C   s   d S r   r   r=   r   r   r   r>   �  s    z"verifyQrcodeWithE2EE_args.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   �  s   �z6verifyQrcodeWithE2EE_args.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   �  s    �z"verifyQrcodeWithE2EE_args.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   �  s    z verifyQrcodeWithE2EE_args.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   �  s    z verifyQrcodeWithE2EE_args.__ne__)NNNNNN)rJ   rU   rV   rW   r   r$   ro   r`   rm   r   r   r1   r3   r>   rN   rR   rS   r   r   r   r   r�   h  s"   �
- r�   c                   @   sn   e Zd ZdZdejdddfdejdedgdffZddd	�Z	d
d� Z
dd� Zdd� Zdd� Zdd� Zdd� ZdS )�verifyQrcodeWithE2EE_resultrY   r   rZ   r   Nr   r[   c                 C   s   || _ || _d S r   r\   r]   r   r   r   r   �  s    z$verifyQrcodeWithE2EE_result.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�tjd dkr�|�� �d�n|�� | _q�|�|� n>|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S r^   r_   r,   r   r   r   r1   �  s&    "

(


z verifyQrcodeWithE2EE_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  | jd k	r�|�dtjd� | j�|� |��  |��  |��  d S )Nr�   rZ   r   r   r   r[   r   rb   r;   r   r   r   r3   
  s    

&
z!verifyQrcodeWithE2EE_result.writec                 C   s   d S r   r   r=   r   r   r   r>     s    z$verifyQrcodeWithE2EE_result.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD     s   �z8verifyQrcodeWithE2EE_result.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN     s    �z$verifyQrcodeWithE2EE_result.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   "  s    z"verifyQrcodeWithE2EE_result.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   %  s    z"verifyQrcodeWithE2EE_result.__ne__)NNrc   r   r   r   r   r�   �  s   �
r�   c                   @   sd   e Zd ZdddejdddffZddd�Zdd� Zdd	� Zd
d� Z	dd� Z
dd� Zdd� Zdd� ZdS )�getRSAKeyInfo_argsNr   �providerc                 C   s
   || _ d S r   �r�   )r   r�   r   r   r   r   0  s    zgetRSAKeyInfo_args.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkrz|tjkrn|�� | _	q�|�
|� n
|�
|� |��  q6|��  d S r�   )r   r   r    r!   r"   r   r#   ro   rp   r�   r)   r*   r+   r,   r   r   r   r1   3  s    



zgetRSAKeyInfo_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� |�
| j� |��  |��  |��  d S )Nr�   r�   r   )r2   r   r   r3   r    r4   r�   r5   r   ro   rr   r8   r9   r:   r;   r   r   r   r3   F  s    

zgetRSAKeyInfo_args.writec                 C   s   d S r   r   r=   r   r   r   r>   R  s    zgetRSAKeyInfo_args.validatec                 C   s   d}|d t | j�A }|S �N�   �   )�hashr�   �r   rC   r   r   r   �__hash__U  s    zgetRSAKeyInfo_args.__hash__c                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   [  s   �z/getRSAKeyInfo_args.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   Z  s    �zgetRSAKeyInfo_args.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   _  s    zgetRSAKeyInfo_args.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   b  s    zgetRSAKeyInfo_args.__ne__)N)rJ   rU   rV   r   ro   r   r   r1   r3   r>   r�   rN   rR   rS   r   r   r   r   r�   (  s   �
r�   c                   @   sv   e Zd ZdejdedgdfdejdedgdffZddd�Zdd	� Z	d
d� Z
dd� Zdd� Zdd� Zdd� Zdd� ZdS )�getRSAKeyInfo_resultr   rZ   Nr   r[   c                 C   s   || _ || _d S r   r\   r]   r   r   r   r   j  s    zgetRSAKeyInfo_result.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n>|dkr�|tjkr�t� | _| j�
|� q�|�|� n
|�|� |��  q6|��  d S r�   )r   r   r    r!   r"   r   r#   r`   �RSAKeyrZ   r1   r)   ra   r[   r*   r+   r,   r   r   r   r1   m  s(    




zgetRSAKeyInfo_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr�   rZ   r   r[   r   r�   r;   r   r   r   r3   �  s    


zgetRSAKeyInfo_result.writec                 C   s   d S r   r   r=   r   r   r   r>   �  s    zgetRSAKeyInfo_result.validatec                 C   s,   d}|d t | j�A }|d t | j�A }|S r�   �r�   rZ   r[   r�   r   r   r   r�   �  s    zgetRSAKeyInfo_result.__hash__c                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   �  s   �z1getRSAKeyInfo_result.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   �  s    �zgetRSAKeyInfo_result.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   �  s    zgetRSAKeyInfo_result.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   �  s    zgetRSAKeyInfo_result.__ne__)NN)rJ   rU   rV   r   r`   r�   ra   r   r   r1   r3   r>   r�   rN   rR   rS   r   r   r   r   r�   e  s   �
r�   c                   @   sj   e Zd ZdZddddejdddffZddd�Zdd	� Zd
d� Z	dd� Z
dd� Zdd� Zdd� Zdd� ZdS )�$loginWithVerifierForCertificate_argsz%
    Attributes:
     - verifier
    Nr   re   r   c                 C   s
   || _ d S r   )re   �r   re   r   r   r   r   �  s    z-loginWithVerifierForCertificate_args.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkr�tj	d dkr~|�
� �d�n|�
� | _q�|�|� n
|�|� |��  q6|��  d S )Nr   r   r   r   )r   r   r    r!   r"   r   r#   r$   r%   r&   r'   r(   re   r)   r*   r+   r,   r   r   r   r1   �  s    

(

z)loginWithVerifierForCertificate_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  |��  |��  d S )Nr�   re   r   r   r   r   )r2   r   r   r3   r    r4   re   r5   r   r$   r6   r%   r&   r7   r8   r9   r:   r;   r   r   r   r3   �  s    

&z*loginWithVerifierForCertificate_args.writec                 C   s   d S r   r   r=   r   r   r   r>   �  s    z-loginWithVerifierForCertificate_args.validatec                 C   s   d}|d t t| j��A }|S r�   )r�   Zmake_hashablere   r�   r   r   r   r�   �  s    z-loginWithVerifierForCertificate_args.__hash__c                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   �  s   �zAloginWithVerifierForCertificate_args.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   �  s    �z-loginWithVerifierForCertificate_args.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   �  s    z+loginWithVerifierForCertificate_args.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   �  s    z+loginWithVerifierForCertificate_args.__ne__)N)rJ   rU   rV   rW   r   r$   r   r   r1   r3   r>   r�   rN   rR   rS   r   r   r   r   r�   �  s   �
r�   c                   @   sz   e Zd ZdZdejdedgdfdejdedgdffZddd�Z	d	d
� Z
dd� Zdd� Zdd� Zdd� Zdd� Zdd� ZdS )�&loginWithVerifierForCertificate_result�-
    Attributes:
     - success
     - e
    r   rZ   Nr   r[   c                 C   s   || _ || _d S r   r\   r]   r   r   r   r   �  s    z/loginWithVerifierForCertificate_result.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n>|dkr�|tjkr�t� | _| j�
|� q�|�|� n
|�|� |��  q6|��  d S r�   �r   r   r    r!   r"   r   r#   r`   �LoginResultrZ   r1   r)   ra   r[   r*   r+   r,   r   r   r   r1   �  s(    




z+loginWithVerifierForCertificate_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr�   rZ   r   r[   r   r�   r;   r   r   r   r3     s    


z,loginWithVerifierForCertificate_result.writec                 C   s   d S r   r   r=   r   r   r   r>   %  s    z/loginWithVerifierForCertificate_result.validatec                 C   s,   d}|d t | j�A }|d t | j�A }|S r�   r�   r�   r   r   r   r�   (  s    z/loginWithVerifierForCertificate_result.__hash__c                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   /  s   �zCloginWithVerifierForCertificate_result.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   .  s    �z/loginWithVerifierForCertificate_result.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   3  s    z-loginWithVerifierForCertificate_result.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   6  s    z-loginWithVerifierForCertificate_result.__ne__)NN�rJ   rU   rV   rW   r   r`   r�   ra   r   r   r1   r3   r>   r�   rN   rR   rS   r   r   r   r   r�   �  s   �
r�   c                	   @   s�   e Zd ZdddejdddfdejdddfdejdddffZdd	d
�Zdd� Zdd� Z	dd� Z
dd� Zdd� Zdd� Zdd� ZdS )�getAuthQrcode_argsNr   �keepLoggedInr   �
systemNamer   r   �callbackUrlc                 C   s   || _ || _|| _d S r   )r�   r�   r�   )r   r�   r�   r�   r   r   r   r   C  s    zgetAuthQrcode_args.__init__c                 C   s  |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr`�q|dkr�|t	jkr~|�� | _n
|�|� nz|dkr�|t	jkr�tjd dkr�|�� �d�n|�� | _n
|�|� n4|dkr�|t	jkr�|�� | _n
|�|� n
|�|� |��  qD|��  d S )Nr   r   r   r   r   )r   r   r   r   r   r   r    r!   r"   r   r#   r�   r�   r�   r)   r$   r%   r&   r'   r(   r�   r�   r*   r+   r,   r   r   r   r1   H  s,    "


(


zgetAuthQrcode_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� |�
| j� |��  | jd k	r�|�dtjd� |�tjd dkr�| j�d�n| j� |��  | jd k	r�|�dtj	d	� |�
| j� |��  |��  |��  d S )
Nr�   r�   r   r�   r   r   r   r�   r   )r2   r   r   r3   r    r4   r�   r5   r   r�   r�   r8   r�   r$   r6   r%   r&   r7   r�   r9   r:   r;   r   r   r   r3   e  s$    


&
zgetAuthQrcode_args.writec                 C   s   d S r   r   r=   r   r   r   r>   y  s    zgetAuthQrcode_args.validatec                 C   s>   d}|d t | j�A }|d t | j�A }|d t | j�A }|S r�   )r�   r�   r�   r�   r�   r   r   r   r�   |  s
    zgetAuthQrcode_args.__hash__c                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   �  s   �z/getAuthQrcode_args.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   �  s    �zgetAuthQrcode_args.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   �  s    zgetAuthQrcode_args.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   �  s    zgetAuthQrcode_args.__ne__)NNN)rJ   rU   rV   r   r�   r$   r   r   r1   r3   r>   r�   rN   rR   rS   r   r   r   r   r�   9  s   �
r�   c                   @   sv   e Zd ZdejdedgdfdejdedgdffZddd�Zdd	� Z	d
d� Z
dd� Zdd� Zdd� Zdd� Zdd� ZdS )�getAuthQrcode_resultr   rZ   Nr   r[   c                 C   s   || _ || _d S r   r\   r]   r   r   r   r   �  s    zgetAuthQrcode_result.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n>|dkr�|tjkr�t� | _| j�
|� q�|�|� n
|�|� |��  q6|��  d S r�   )r   r   r    r!   r"   r   r#   r`   �
AuthQrcoderZ   r1   r)   ra   r[   r*   r+   r,   r   r   r   r1   �  s(    




zgetAuthQrcode_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr�   rZ   r   r[   r   r�   r;   r   r   r   r3   �  s    


zgetAuthQrcode_result.writec                 C   s   d S r   r   r=   r   r   r   r>   �  s    zgetAuthQrcode_result.validatec                 C   s,   d}|d t | j�A }|d t | j�A }|S r�   r�   r�   r   r   r   r�   �  s    zgetAuthQrcode_result.__hash__c                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   �  s   �z1getAuthQrcode_result.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   �  s    �zgetAuthQrcode_result.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   �  s    zgetAuthQrcode_result.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   �  s    zgetAuthQrcode_result.__ne__)NN)rJ   rU   rV   r   r`   r�   ra   r   r   r1   r3   r>   r�   rN   rR   rS   r   r   r   r   r�   �  s   �
r�   c                   @   s    e Zd ZdZdd� Zdd� ZdS )�logoutZ_argsr   c                 C   sr   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrRqfn
|�|� |��  q6|�	�  d S r   )
r   r   r    r!   r"   r   r#   r)   r*   r+   r,   r   r   r   r1   �  s    


zlogoutZ_args.readc                 C   sR   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� |��  |��  d S )Nr�   )r2   r   r   r3   r    r4   r9   r:   r;   r   r   r   r3   �  s    
zlogoutZ_args.writeN)rJ   rU   rV   r   r1   r3   r   r   r   r   r�   �  s   r�   c                   @   sB   e Zd ZdZddejdedgdffZddd�Zdd� Z	d	d
� Z
dS )�logoutZ_resultz
    Attributes:
     - e
    Nr   r[   c                 C   s
   || _ d S r   ru   rv   r   r   r   r   �  s    zlogoutZ_result.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n
|�|� |��  q6|��  d S rw   )r   r   r    r!   r"   r   r#   r`   ra   r[   r1   r)   r*   r+   r,   r   r   r   r1      s    



zlogoutZ_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr�   r[   r   ry   r;   r   r   r   r3     s    

zlogoutZ_result.write)N)rJ   rU   rV   rW   r   r`   ra   r   r   r1   r3   r   r   r   r   r�   �  s   �
r�   c                   @   sl   e Zd ZdZdddejdedgdffZddd�Zdd� Z	d	d
� Z
dd� Zdd� Zdd� Zdd� Zdd� ZdS )�loginZ_argsz)
    Attributes:
     - loginRequest
    Nr   �loginRequestc                 C   s
   || _ d S r   )r�   )r   r�   r   r   r   r   ,  s    zloginZ_args.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n
|�|� |��  q6|��  d S r�   )r   r   r    r!   r"   r   r#   r`   �LoginRequestr�   r1   r)   r*   r+   r,   r   r   r   r1   /  s    



zloginZ_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr�   r�   r   )r2   r   r   r3   r    r4   r�   r5   r   r`   r8   r9   r:   r;   r   r   r   r3   C  s    

zloginZ_args.writec                 C   s   d S r   r   r=   r   r   r   r>   O  s    zloginZ_args.validatec                 C   s   d}|d t | j�A }|S r�   )r�   r�   r�   r   r   r   r�   R  s    zloginZ_args.__hash__c                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   X  s   �z(loginZ_args.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   W  s    �zloginZ_args.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   \  s    zloginZ_args.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   _  s    zloginZ_args.__ne__)N)rJ   rU   rV   rW   r   r`   r�   r   r   r1   r3   r>   r�   rN   rR   rS   r   r   r   r   r�   !  s   �
r�   c                   @   sz   e Zd ZdZdejdedgdfdejdedgdffZddd�Z	d	d
� Z
dd� Zdd� Zdd� Zdd� Zdd� Zdd� ZdS )�loginZ_resultr�   r   rZ   Nr   r[   c                 C   s   || _ || _d S r   r\   r]   r   r   r   r   m  s    zloginZ_result.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n>|dkr�|tjkr�t� | _| j�
|� q�|�|� n
|�|� |��  q6|��  d S r�   r�   r,   r   r   r   r1   q  s(    




zloginZ_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr�   rZ   r   r[   r   r�   r;   r   r   r   r3   �  s    


zloginZ_result.writec                 C   s   d S r   r   r=   r   r   r   r>   �  s    zloginZ_result.validatec                 C   s,   d}|d t | j�A }|d t | j�A }|S r�   r�   r�   r   r   r   r�   �  s    zloginZ_result.__hash__c                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rD   �  s   �z*loginZ_result.__repr__.<locals>.<listcomp>rE   rF   rG   rL   r   r   r   rN   �  s    �zloginZ_result.__repr__c                 C   s   t || j�o| j|jkS r   rO   rP   r   r   r   rR   �  s    zloginZ_result.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   �  s    zloginZ_result.__ne__)NNr�   r   r   r   r   r�   b  s   �
r�   c                   @   s�   e Zd Zd<dd�Zdd� Zdd� Zdd	� Zd
d� Zdd� Zdd� Z	dd� Z
dd� Zdd� Zdd� Zdd� Zdd� Zdd� Zdd� Zd d!� Zd"d#� Zd$d%� Zd&d'� Zd(d)� Zd*d+� Zd,d-� Zd.d/� Zd0d1� Zd2d3� Zd4d5� Zd6d7� Zd8d9� Zd:d;� ZdS )=�ClientNc                 C   s$   | | _ | _|d k	r|| _d| _d S )Nr   )�_iprot�_oprot�_seqid)r   r-   r<   r   r   r   r   �  s    zClient.__init__c                 C   s   | � |||� | �� S )zg
        Parameters:
         - countryCode
         - phoneNumber
         - countryCodeHint

        )�send_normalizePhoneNumber�recv_normalizePhoneNumberr   r   r   r   �normalizePhoneNumber�  s    zClient.normalizePhoneNumberc                 C   sR   | j �dtj| j� t� }||_||_||_|�	| j � | j �
�  | j j��  d S )Nr�   )r�   �writeMessageBeginr   �CALLr�   r   r   r   r   r3   �writeMessageEndr   �flush)r   r   r   r   �argsr   r   r   r�   �  s    
z Client.send_normalizePhoneNumberc                 C   s�   | j }|�� \}}}|tjkr:t� }|�|� |��  |�t� }|�|� |��  |jd k	rb|jS |j	d k	rr|j	�ttj
d��d S )Nz+normalizePhoneNumber failed: unknown result)r�   �readMessageBeginr   �	EXCEPTIONr   r1   �readMessageEndrX   rZ   r[   �MISSING_RESULT�r   r-   r.   �mtype�rseqid�x�resultr   r   r   r�   �  s    




z Client.recv_normalizePhoneNumberc                 C   s   | � |||||� | ��  dS )z�
        Parameters:
         - verifier
         - publicKey
         - encryptedKeyChain
         - hashKeyChain
         - errorCode

        N)�send_respondE2EELoginRequest�recv_respondE2EELoginRequestrl   r   r   r   �respondE2EELoginRequest�  s    
zClient.respondE2EELoginRequestc                 C   s^   | j �dtj| j� t� }||_||_||_||_	||_
|�| j � | j ��  | j j��  d S )Nr�   )r�   r�   r   r�   r�   rd   re   rf   rg   ri   rk   r3   r�   r   r�   )r   re   rf   rg   ri   rk   r�   r   r   r   r�   �  s    
z#Client.send_respondE2EELoginRequestc                 C   sf   | j }|�� \}}}|tjkr:t� }|�|� |��  |�t� }|�|� |��  |jd k	rb|j�d S r   )	r�   r�   r   r�   r   r1   r�   rs   r[   r�   r   r   r   r�   �  s    



z#Client.recv_respondE2EELoginRequestc                 C   s   | � ||� | ��  dS )zS
        Parameters:
         - authSessionId
         - identityProvider

        N)�send_getAuthRSAKey�recv_getAuthRSAKeyr~   r   r   r   �getAuthRSAKey  s    zClient.getAuthRSAKeyc                 C   sL   | j �dtj| j� t� }||_||_|�| j � | j �	�  | j j
��  d S )Nr�   )r�   r�   r   r�   r�   r{   r|   r}   r3   r�   r   r�   )r   r|   r}   r�   r   r   r   r�     s    
zClient.send_getAuthRSAKeyc                 C   sf   | j }|�� \}}}|tjkr:t� }|�|� |��  |�t� }|�|� |��  |jd k	rb|j�d S r   )	r�   r�   r   r�   r   r1   r�   r   r[   r�   r   r   r   r�     s    



zClient.recv_getAuthRSAKeyc                 C   s   | � ||� | �� S )zJ
        Parameters:
         - verifier
         - deviceSecret

        )�send_confirmE2EELogin�recv_confirmE2EELoginr�   r   r   r   �confirmE2EELogin&  s    zClient.confirmE2EELoginc                 C   sL   | j �dtj| j� t� }||_||_|�| j � | j �	�  | j j
��  d S )Nr�   )r�   r�   r   r�   r�   r�   re   r�   r3   r�   r   r�   )r   re   r�   r�   r   r   r   r�   0  s    
zClient.send_confirmE2EELoginc                 C   s�   | j }|�� \}}}|tjkr:t� }|�|� |��  |�t� }|�|� |��  |jd k	rb|jS |j	d k	rr|j	�ttj
d��d S )Nz'confirmE2EELogin failed: unknown result)r�   r�   r   r�   r   r1   r�   r�   rZ   r[   r�   r�   r   r   r   r�   9  s    




zClient.recv_confirmE2EELoginc                 C   s   | � |� | �� S )z1
        Parameters:
         - enforce

        )�*send_issueTokenForAccountMigrationSettings�*recv_issueTokenForAccountMigrationSettingsr�   r   r   r   �%issueTokenForAccountMigrationSettingsJ  s    
z,Client.issueTokenForAccountMigrationSettingsc                 C   sF   | j �dtj| j� t� }||_|�| j � | j ��  | j j	�
�  d S )Nr�   )r�   r�   r   r�   r�   r�   r�   r3   r�   r   r�   )r   r�   r�   r   r   r   r�   S  s    
z1Client.send_issueTokenForAccountMigrationSettingsc                 C   s�   | j }|�� \}}}|tjkr:t� }|�|� |��  |�t� }|�|� |��  |jd k	rb|jS |j	d k	rr|j	�ttj
d��d S )Nz<issueTokenForAccountMigrationSettings failed: unknown result)r�   r�   r   r�   r   r1   r�   r�   rZ   r[   r�   r�   r   r   r   r�   [  s    




z1Client.recv_issueTokenForAccountMigrationSettingsc                 C   s   | � |� | �� S )z<
        Parameters:
         - migrationSessionId

        )�"send_issueTokenForAccountMigration�"recv_issueTokenForAccountMigrationr�   r   r   r   �issueTokenForAccountMigrationl  s    
z$Client.issueTokenForAccountMigrationc                 C   sF   | j �dtj| j� t� }||_|�| j � | j ��  | j j	�
�  d S )Nr�   )r�   r�   r   r�   r�   r�   r�   r3   r�   r   r�   )r   r�   r�   r   r   r   r�   u  s    
z)Client.send_issueTokenForAccountMigrationc                 C   s�   | j }|�� \}}}|tjkr:t� }|�|� |��  |�t� }|�|� |��  |jd k	rb|jS |j	d k	rr|j	�ttj
d��d S )Nz4issueTokenForAccountMigration failed: unknown result)r�   r�   r   r�   r   r1   r�   r�   rZ   r[   r�   r�   r   r   r   r�   }  s    




z)Client.recv_issueTokenForAccountMigrationc                 C   s   | � ||||||� | �� S )z�
        Parameters:
         - verifier
         - pinCode
         - errorCode
         - publicKey
         - encryptedKeyChain
         - hashKeyChain

        )�send_verifyQrcodeWithE2EE�recv_verifyQrcodeWithE2EEr�   r   r   r   �verifyQrcodeWithE2EE�  s    zClient.verifyQrcodeWithE2EEc                 C   sd   | j �dtj| j� t� }||_||_||_||_	||_
||_|�| j � | j ��  | j j��  d S )Nr�   )r�   r�   r   r�   r�   r�   re   r�   rk   rf   rg   ri   r3   r�   r   r�   )r   re   r�   rk   rf   rg   ri   r�   r   r   r   r�   �  s    
z Client.send_verifyQrcodeWithE2EEc                 C   s�   | j }|�� \}}}|tjkr:t� }|�|� |��  |�t� }|�|� |��  |jd k	rb|jS |j	d k	rr|j	�ttj
d��d S )Nz+verifyQrcodeWithE2EE failed: unknown result)r�   r�   r   r�   r   r1   r�   r�   rZ   r[   r�   r�   r   r   r   r�   �  s    




z Client.recv_verifyQrcodeWithE2EEc                 C   s�   | j �dtj| j� t|d��| j � | j ��  | j j�	�  | j
}|�� \}}}|tjkrvt� }|�|� |��  |�t� }|�|� |��  |jd k	r�|jS |jd k	r�t|j��ttjd��d S )N�getRSAKeyInfor�   z$getRSAKeyInfo failed: unknown result)r�   r�   r   r�   r�   r�   r3   r�   r   r�   r�   r�   r�   r   r1   r�   r�   rZ   r[   �	Exceptionr�   )r   r�   r-   r.   r�   r�   r�   r�   r   r   r   r�   �  s&    






zClient.getRSAKeyInfoc                 C   s   | � |� | �� S )z1
        Parameters:
         - verifier
        )�$send_loginWithVerifierForCertificate�$recv_loginWithVerifierForCertificater�   r   r   r   �loginWithVerifierForCertificate�  s    
z&Client.loginWithVerifierForCertificatec                 C   sN   | j �dtj| j� t� }||_|�| j � |��  | j �	�  | j j
��  d S )Nr�   )r�   r�   r   r�   r�   r�   re   r3   r�   r�   r   r�   )r   re   r�   r   r   r   r�   �  s    
z+Client.send_loginWithVerifierForCertificatec                 C   s�   | j }|�� \}}}|tjkr:t� }|�|� |��  |�t� }|�|� |��  |jd k	rb|jS |j	d k	rr|j	�ttj
d��d S )Nz6loginWithVerifierForCertificate failed: unknown result)r�   r�   r   r�   r   r1   r�   r�   rZ   r[   r�   r�   r   r   r   r�   �  s    




z+Client.recv_loginWithVerifierForCertificatec                 C   s�   | j �dtj| j� t� }||_||_||_|�	| j � | j �
�  | j j��  | j}|�� \}}}|tjkr�t� }	|	�|� |��  |	�t� }
|
�|� |��  |
jd k	r�|
jS |
jd k	r�|
j�ttjd��d S )N�getAuthQrcodez$getAuthQrcode failed: unknown result)r�   r�   r   r�   r�   r�   r�   r�   r�   r3   r�   r   r�   r�   r�   r�   r   r1   r�   r�   rZ   r[   r�   )r   r�   r�   r�   r�   r-   r.   r�   r�   r�   r�   r   r   r   r�   �  s.    





zClient.getAuthQrcodec                 C   s�   | j �dtj| j� t� }|�| j � | j ��  | j j�	�  | j
}|�� \}}}|tjkrvt� }|�|� |��  |�t� }|�|� |��  |jd k	r�|j�d S )N�logoutZ)r�   r�   r   r�   r�   r�   r3   r�   r   r�   r�   r�   r�   r   r1   r�   r�   r[   )r   r�   r-   r.   r�   r�   r�   r�   r   r   r   r�   
  s$    




zClient.logoutZc           	      C   s�   | j �dtj| j� t� }||_|�| j � | j ��  | j j	�
�  | j}|�� \}}}|tjkr|t� }|�|� |��  |�t� }|�|� |��  |jd k	r�|jS |jd k	r�|j�ttjd��d S )N�loginZzloginZ failed: unknown result)r�   r�   r   r�   r�   r�   r�   r3   r�   r   r�   r�   r�   r�   r   r1   r�   r�   rZ   r[   r�   )	r   r�   r�   r-   r.   r�   r�   r�   r�   r   r   r   r�     s*    





zClient.loginZ)N) rJ   rU   rV   r   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r   r   r   r   r�   �  s:   


	
				r�   )(�thrift.Thriftr   r   r   r   r   �thrift.protocol.TProtocolr   r%   Zlogging�ttypesr
   �thrift.transportr   �objectr   rX   rd   rs   r{   r   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r�   r   r   r   r   �<module>   s>   TGk<H<GG<H<HyG=EBMUI.AM